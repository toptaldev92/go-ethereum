package api

const Miner_JS = `
web3.extend({
	property: 'miner',
	methods:
	[
		new web3.extend.Method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [web3.extend.formatters.formatInputInt],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'stop',
			call: 'miner_stop',
			params: 1,
			inputFormatter: [web3.extend.formatters.formatInputInt],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'getHashrate',
			call: 'miner_hashrate',
			params: 0,
			inputFormatter: [],
			outputFormatter: web3.extend.utils.toDecimal
		}),
		new web3.extend.Method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1,
			inputFormatter: [web3.extend.utils.formatInputString],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3.extend.utils.formatInputString],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'startAutoDAG',
			call: 'miner_startAutoDAG',
			params: 0,
			inputFormatter: [],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'stopAutoDAG',
			call: 'miner_stopAutoDAG',
			params: 0,
			inputFormatter: [],
			outputFormatter: web3.extend.formatters.formatOutputBool
		}),
		new web3.extend.Method({
			name: 'makeDAG',
			call: 'miner_makeDAG',
			params: 1,
			inputFormatter: [web3.extend.formatters.inputDefaultBlockNumberFormatter],
			outputFormatter: web3.extend.formatters.formatOutputBool
		})
	],
	properties:
	[
		new web3.extend.Property({
			name: 'hashrate',
			getter: 'miner_hashrate',
			outputFormatter: web3.extend.utils.toDecimal
		})
	]
});
`