<!--
order: 1
markdown-link-check-disable
-->

# ¿Qué es empyrean?

`empyrean` es el nombre de la aplicación de Cosmos SDK para el Hub de Cosmos. Viene con 2 opciones principales de entrada:

- `empyreand`: El demonio de empyrean, ejecuta un full-node con la aplicacion `empyrean`.
- `empyreand`: Interface de línea de comandos de empyrean, la cual activa la interacción con en full-node de empyrean.

`empyrean` es construida con el SDK de Cosmos usando los siguientes módulos:

- `x/auth`: Cuentas y firmas.
- `x/bank`: Transferencia de tokens.
- `x/staking`: Lógica para el _staking_.
- `x/mint`: Lógica para la inflacción.
- `x/distribution`: Lógica para la distribución del FEE.
- `x/slashing`: Lógica para la penalización.
- `x/gov`: Lógica para la gobernanza.
- `ibc-go/modules`: Transferencia entre blockchains.
- `x/params`: Controla los parámetros del nivel de la aplicación.
- `x/capability`:
- `x/crisis`:
- `x/evidence`:
- `x/genaccounts`:
- `x/mint`:
- `x/simulation`:
- `x/upgrade`:

Acerca del Hub de Cosmos: El Hub de Cosmos es el primer Hub en ser lanzado en la red de Cosmos. El propósito del Hub es facilitar las transferencias entre cadenas de bloques. Si una cadena de bloques se conecta a un Hub a través de IBC, automáticamente obtiene acceso a todas las otras cadenas de bloques que están conectadas a ella. El Hub del Cosmos es una cadena de participación pública. Su token de participación es llamado Atom.

Siguiente, aprenda como [instalar empyrean](./installation.md)

<!-- markdown-link-check-enable -->