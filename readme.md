# Lightmos Hub (empyrean)

![Banner!](https://miro.medium.com/max/2000/1*DHtmSfS_Efvuq8n2LAnhkA.png)

[![Project Status: Active -- The project has reached a stable, usable state and is being actively
developed.](https://img.shields.io/badge/repo%20status-Active-green.svg)](https://www.repostatus.org/#active)

The Lightmos Hub is a interconnected blockchains that comprise the Cosmos Network.

<br/>

## 🤔 — Why should you be interested in the Lightmos Hub

___

The Lightmos Hub is built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) and compiled to a binary called `empyreand` (empyrean Daemon). The Cosmos Hub and other fully sovereign Cosmos SDK blockchains interact with one another using a protocol called [IBC](https://github.com/cosmos/ibc) that enables Inter-Blockchain Communication. In order to understand what the Cosmos Hub is you can read this [introductory explanation](https://hub.cosmos.network/main/hub-overview/overview.html).

<br/>

## ⚡ — Documentation & Introduction

___

Cosmos Hub is a blockchain network that operates on Proof-of-Stake consensus. You can find an introduction to the Cosmos Hub and how to use the `empyreand` binary as a delegator, validator or node operator as well as how governance on the Cosmos Hub works in the [documentation](https://hub.cosmos.network/main/hub-overview/overview.html).

Alternatively, whether you're new to blockchain technology or interested in getting involved, the Cosmos Network [Course](https://tutorials.cosmos.network/academy/0-welcome/) will guide you through everything. The course walks you through the basics of blockchain technology, to staking, setting up your own node, and beyond.

<br/>


## 🗣️ — Validators

___

If you want to participate and help secure Cosmos Hub, check out becoming a validator. Information on what a validator is and how to participate as one can be found at the [Validator FAQ](https://hub.cosmos.network/main/validators/validator-faq.html#). If you're running a validator node on the Cosmos Hub, reach out to a Janitor on the [Cosmos Developers Discord](https://discord.gg/cosmosnetwork) to join the `#cosmos-hub-validators-verified` channel.

<br/>

## 👥 — Delegators

___

If you still want to participate on the Cosmos Hub, check out becoming a delegator. Information on what a delegator is and how to participate as one can be found at the [Delegator FAQ](https://hub.cosmos.network/main/delegators/delegator-faq.html).

<br/>


# empyrean
**empyrean** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/empyrean@latest! | sudo bash
```
`username/empyrean` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
