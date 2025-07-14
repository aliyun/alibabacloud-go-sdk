// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SmartqAuthTransferResponseBody
	GetRequestId() *string
	SetResult(v bool) *SmartqAuthTransferResponseBody
	GetResult() *bool
	SetSuccess(v bool) *SmartqAuthTransferResponseBody
	GetSuccess() *bool
}

type SmartqAuthTransferResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1*****************5DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API execution result. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SmartqAuthTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthTransferResponseBody) GoString() string {
	return s.String()
}

func (s *SmartqAuthTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmartqAuthTransferResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SmartqAuthTransferResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SmartqAuthTransferResponseBody) SetRequestId(v string) *SmartqAuthTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartqAuthTransferResponseBody) SetResult(v bool) *SmartqAuthTransferResponseBody {
	s.Result = &v
	return s
}

func (s *SmartqAuthTransferResponseBody) SetSuccess(v bool) *SmartqAuthTransferResponseBody {
	s.Success = &v
	return s
}

func (s *SmartqAuthTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
