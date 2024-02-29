SELECT 
    B.CadastroId,
    A.Descricao AS DiaSemana,
    COALESCE(SUM(C.Horas), 0) AS TotalHoras
FROM 
    DiasSemana A
CROSS JOIN 
    (SELECT DISTINCT CadastroId FROM HorariosTrabalho) AS B
LEFT JOIN 
    HorariosTrabalho C ON A.DiaId = C.DiaId AND B.CadastroId = C.CadastroId
GROUP BY 
    B.CadastroId,
    A.Descricao
ORDER BY 
    B.CadastroId,
    A.DiaId;
