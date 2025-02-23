# Arrays et Slices en Go

## Introduction
En Go, un **array** (tableau) est une structure de données qui contient un nombre fixe d'éléments du même type. Contrairement aux slices, la taille d'un array est définie lors de sa déclaration et ne peut pas être modifiée.

## Déclaration et Initialisation
### Déclaration d'un Array
Un array peut être déclaré en spécifiant son type et sa taille :

```go
var numbers [5]int // Déclare un tableau de 5 entiers
```

### Initialisation d'un Array
On peut initialiser un array lors de sa déclaration :

```go
var numbers = [5]int{1, 2, 3, 4, 5}
```

Ou utiliser l'inférence de longueur :

```go
numbers := [...]int{1, 2, 3, 4, 5} // La taille est déterminée automatiquement
```

## Accès aux éléments
Les éléments d'un array sont accessibles via des indices (commençant à 0) :

```go
fmt.Println(numbers[0]) // Affiche le premier élément
```

Modification d'un élément :

```go
numbers[1] = 10 // Modifie le deuxième élément
```

## Parcourir un Array
On peut parcourir un array en utilisant une boucle `for` :

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

Ou avec `range` :

```go
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Complexité des opérations
Les opérations sur les arrays ont différentes complexités :
- **Accès à un élément** : `O(1)`
- **Modification d'un élément** : `O(1)`
- **Parcours** : `O(n)`
- **Copie** : `O(n)`

Les arrays sont efficaces pour l'accès direct, mais peu flexibles en raison de leur taille fixe.

## Slices : Une alternative flexible
Un **slice** est une abstraction des arrays qui permet une taille dynamique. Un slice est défini à partir d'un array sous-jacent et peut être redimensionné avec `append`.

### Déclaration et Initialisation d'un Slice

```go
var slice []int // Déclare un slice vide
slice = []int{1, 2, 3} // Initialisation
```

Utilisation de `append` pour ajouter des éléments dynamiquement :

```go
slice = append(slice, 4, 5)
fmt.Println(slice) // [1, 2, 3, 4, 5]
```

### Opérations courantes sur les slices
- **Capacité (`cap()`)** : Retourne la capacité du slice (taille de l'array sous-jacent)
- **Longueur (`len()`)** : Retourne le nombre d'éléments dans le slice
- **Sous-slice (`slice[start:end]`)** : Extrait une partie du slice

```go
subSlice := slice[1:3] // Contient [2, 3]
```

## Copie et Références
En Go, les slices sont des références à des arrays sous-jacents. Modifier un élément d'un slice affecte l'array d'origine :

```go
slice := []int{1, 2, 3}
subSlice := slice[:2]
subSlice[0] = 100
fmt.Println(slice) // [100, 2, 3]
```

Pour copier un slice indépendamment, utilisez `copy` :

```go
copySlice := make([]int, len(slice))
copy(copySlice, slice)
```

## Comparaison Arrays vs Slices
| Fonctionnalité | Arrays | Slices |
|--------------|--------|--------|
| Taille fixe | ✅ | ❌ |
| Dynamique | ❌ | ✅ |
| Utilisation de `append` | ❌ | ✅ |
| Allocation mémoire flexible | ❌ | ✅ |

## Conclusion
Les arrays en Go sont utiles pour stocker un nombre fixe d'éléments. Cependant, dans la majorité des cas, les **slices** sont préférables en raison de leur flexibilité. Comprendre les arrays est toutefois essentiel car les slices reposent sur leur fonctionnement interne.

---

### Ressources supplémentaires
- [Documentation officielle Go sur les arrays](https://go.dev/doc/effective_go#arrays)
- [Go by Example: Arrays](https://gobyexample.com/arrays)

