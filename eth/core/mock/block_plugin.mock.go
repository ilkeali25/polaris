// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"pkg.berachain.dev/polaris/eth/core"
	"sync"
)

// Ensure, that BlockPluginMock does implement core.BlockPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.BlockPlugin = &BlockPluginMock{}

// BlockPluginMock is a mock implementation of core.BlockPlugin.
//
//	func TestSomethingThatUsesBlockPlugin(t *testing.T) {
//
//		// make and configure a mocked core.BlockPlugin
//		mockedBlockPlugin := &BlockPluginMock{
//			BaseFeeFunc: func() uint64 {
//				panic("mock out the BaseFee method")
//			},
//			GetHeaderByNumberFunc: func(n int64) (*types.Header, error) {
//				panic("mock out the GetHeaderByNumber method")
//			},
//			GetNewBlockMetadataFunc: func(n int64) (common.Address, uint64) {
//				panic("mock out the GetNewBlockMetadata method")
//			},
//			PrepareFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Prepare method")
//			},
//			SetHeaderByNumberFunc: func(n int64, header *types.Header) error {
//				panic("mock out the SetHeaderByNumber method")
//			},
//		}
//
//		// use mockedBlockPlugin in code that requires core.BlockPlugin
//		// and then make assertions.
//
//	}
type BlockPluginMock struct {
	// BaseFeeFunc mocks the BaseFee method.
	BaseFeeFunc func() uint64

	// GetHeaderByNumberFunc mocks the GetHeaderByNumber method.
	GetHeaderByNumberFunc func(n int64) (*types.Header, error)

	// GetNewBlockMetadataFunc mocks the GetNewBlockMetadata method.
	GetNewBlockMetadataFunc func(n int64) (common.Address, uint64)

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(contextMoqParam context.Context)

	// SetHeaderByNumberFunc mocks the SetHeaderByNumber method.
	SetHeaderByNumberFunc func(n int64, header *types.Header) error

	// calls tracks calls to the methods.
	calls struct {
		// BaseFee holds details about calls to the BaseFee method.
		BaseFee []struct {
		}
		// GetHeaderByNumber holds details about calls to the GetHeaderByNumber method.
		GetHeaderByNumber []struct {
			// N is the n argument value.
			N int64
		}
		// GetNewBlockMetadata holds details about calls to the GetNewBlockMetadata method.
		GetNewBlockMetadata []struct {
			// N is the n argument value.
			N int64
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// SetHeaderByNumber holds details about calls to the SetHeaderByNumber method.
		SetHeaderByNumber []struct {
			// N is the n argument value.
			N int64
			// Header is the header argument value.
			Header *types.Header
		}
	}
	lockBaseFee             sync.RWMutex
	lockGetHeaderByNumber   sync.RWMutex
	lockGetNewBlockMetadata sync.RWMutex
	lockPrepare             sync.RWMutex
	lockSetHeaderByNumber   sync.RWMutex
}

// BaseFee calls BaseFeeFunc.
func (mock *BlockPluginMock) BaseFee() uint64 {
	if mock.BaseFeeFunc == nil {
		panic("BlockPluginMock.BaseFeeFunc: method is nil but BlockPlugin.BaseFee was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBaseFee.Lock()
	mock.calls.BaseFee = append(mock.calls.BaseFee, callInfo)
	mock.lockBaseFee.Unlock()
	return mock.BaseFeeFunc()
}

// BaseFeeCalls gets all the calls that were made to BaseFee.
// Check the length with:
//
//	len(mockedBlockPlugin.BaseFeeCalls())
func (mock *BlockPluginMock) BaseFeeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBaseFee.RLock()
	calls = mock.calls.BaseFee
	mock.lockBaseFee.RUnlock()
	return calls
}

// GetHeaderByNumber calls GetHeaderByNumberFunc.
func (mock *BlockPluginMock) GetHeaderByNumber(n int64) (*types.Header, error) {
	if mock.GetHeaderByNumberFunc == nil {
		panic("BlockPluginMock.GetHeaderByNumberFunc: method is nil but BlockPlugin.GetHeaderByNumber was just called")
	}
	callInfo := struct {
		N int64
	}{
		N: n,
	}
	mock.lockGetHeaderByNumber.Lock()
	mock.calls.GetHeaderByNumber = append(mock.calls.GetHeaderByNumber, callInfo)
	mock.lockGetHeaderByNumber.Unlock()
	return mock.GetHeaderByNumberFunc(n)
}

// GetHeaderByNumberCalls gets all the calls that were made to GetHeaderByNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.GetHeaderByNumberCalls())
func (mock *BlockPluginMock) GetHeaderByNumberCalls() []struct {
	N int64
} {
	var calls []struct {
		N int64
	}
	mock.lockGetHeaderByNumber.RLock()
	calls = mock.calls.GetHeaderByNumber
	mock.lockGetHeaderByNumber.RUnlock()
	return calls
}

// GetNewBlockMetadata calls GetNewBlockMetadataFunc.
func (mock *BlockPluginMock) GetNewBlockMetadata(n int64) (common.Address, uint64) {
	if mock.GetNewBlockMetadataFunc == nil {
		panic("BlockPluginMock.GetNewBlockMetadataFunc: method is nil but BlockPlugin.GetNewBlockMetadata was just called")
	}
	callInfo := struct {
		N int64
	}{
		N: n,
	}
	mock.lockGetNewBlockMetadata.Lock()
	mock.calls.GetNewBlockMetadata = append(mock.calls.GetNewBlockMetadata, callInfo)
	mock.lockGetNewBlockMetadata.Unlock()
	return mock.GetNewBlockMetadataFunc(n)
}

// GetNewBlockMetadataCalls gets all the calls that were made to GetNewBlockMetadata.
// Check the length with:
//
//	len(mockedBlockPlugin.GetNewBlockMetadataCalls())
func (mock *BlockPluginMock) GetNewBlockMetadataCalls() []struct {
	N int64
} {
	var calls []struct {
		N int64
	}
	mock.lockGetNewBlockMetadata.RLock()
	calls = mock.calls.GetNewBlockMetadata
	mock.lockGetNewBlockMetadata.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *BlockPluginMock) Prepare(contextMoqParam context.Context) {
	if mock.PrepareFunc == nil {
		panic("BlockPluginMock.PrepareFunc: method is nil but BlockPlugin.Prepare was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(contextMoqParam)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedBlockPlugin.PrepareCalls())
func (mock *BlockPluginMock) PrepareCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}

// SetHeaderByNumber calls SetHeaderByNumberFunc.
func (mock *BlockPluginMock) SetHeaderByNumber(n int64, header *types.Header) error {
	if mock.SetHeaderByNumberFunc == nil {
		panic("BlockPluginMock.SetHeaderByNumberFunc: method is nil but BlockPlugin.SetHeaderByNumber was just called")
	}
	callInfo := struct {
		N      int64
		Header *types.Header
	}{
		N:      n,
		Header: header,
	}
	mock.lockSetHeaderByNumber.Lock()
	mock.calls.SetHeaderByNumber = append(mock.calls.SetHeaderByNumber, callInfo)
	mock.lockSetHeaderByNumber.Unlock()
	return mock.SetHeaderByNumberFunc(n, header)
}

// SetHeaderByNumberCalls gets all the calls that were made to SetHeaderByNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.SetHeaderByNumberCalls())
func (mock *BlockPluginMock) SetHeaderByNumberCalls() []struct {
	N      int64
	Header *types.Header
} {
	var calls []struct {
		N      int64
		Header *types.Header
	}
	mock.lockSetHeaderByNumber.RLock()
	calls = mock.calls.SetHeaderByNumber
	mock.lockSetHeaderByNumber.RUnlock()
	return calls
}
