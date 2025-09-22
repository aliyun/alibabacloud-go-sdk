// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFraudResultCallBackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FraudResultCallBackResponseBody
	GetCode() *string
	SetMessage(v string) *FraudResultCallBackResponseBody
	GetMessage() *string
	SetRequestId(v string) *FraudResultCallBackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FraudResultCallBackResponseBody
	GetSuccess() *bool
}

type FraudResultCallBackResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FraudResultCallBackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FraudResultCallBackResponseBody) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponseBody) GetCode() *string {
	return s.Code
}

func (s *FraudResultCallBackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FraudResultCallBackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FraudResultCallBackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FraudResultCallBackResponseBody) SetCode(v string) *FraudResultCallBackResponseBody {
	s.Code = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetMessage(v string) *FraudResultCallBackResponseBody {
	s.Message = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetRequestId(v string) *FraudResultCallBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FraudResultCallBackResponseBody) SetSuccess(v bool) *FraudResultCallBackResponseBody {
	s.Success = &v
	return s
}

func (s *FraudResultCallBackResponseBody) Validate() error {
	return dara.Validate(s)
}
