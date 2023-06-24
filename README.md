
<h1>Program zgadywania liczby w języku Go</h1>

<h2>Zasady gry</h2>
    <ol>
        <li>Program losuje liczbę całkowitą z przedziału od 1 do 100.</li>
        <li>Gracz ma za zadanie zgadnąć wylosowaną liczbę.</li>
        <li>Po każdej próbie zgadnięcia, program informuje gracza, czy liczba jest za mała, za duża lub poprawna.</li>
        <li>Gracz może kontynuować próby zgadywania, aż odgadnie liczbę lub zrezygnuje z gry.</li>
        <li>Po zakończeniu gry, program wyświetla liczbę prób, które gracz podjął, oraz zapisuje wynik do pliku.</li>
    </ol>
 <h2>Zapis wyników</h2>
    <p>Po zakończeniu gry, wyniki są zapisywane do pliku <code>scores.txt</code>. Każdy wynik jest dodawany jako nowa linia w formacie:</p>
    <pre>Data: &lt;data&gt; | Liczba prób: &lt;liczba-prob&gt; | Nazwa: &lt;nick gracza&gt;</pre>

