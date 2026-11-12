// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailCount(v int32) *AttachKeyPairResponseBody
	GetFailCount() *int32
	SetRequestId(v string) *AttachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v []*AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody
	GetResults() []*AttachKeyPairResponseBodyResults
	SetTotalCount(v int32) *AttachKeyPairResponseBody
	GetTotalCount() *int32
}

type AttachKeyPairResponseBody struct {
	// The total number of instances to which the key pair failed to be attached.
	//
	// example:
	//
	// 2
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result list.
	Results []*AttachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of instances to which the key pair is attached.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *AttachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachKeyPairResponseBody) GetResults() []*AttachKeyPairResponseBodyResults {
	return s.Results
}

func (s *AttachKeyPairResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AttachKeyPairResponseBody) SetFailCount(v int32) *AttachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetRequestId(v string) *AttachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetResults(v []*AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *AttachKeyPairResponseBody) SetTotalCount(v int32) *AttachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

func (s *AttachKeyPairResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachKeyPairResponseBodyResults struct {
	// The result code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// aa6e71ddb35c46679bc4753d6219d604
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The result message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the key pair is successfully attached. Valid values:
	//
	// - true: The key pair is successfully attached.
	//
	// - false: The key pair failed to be attached.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachKeyPairResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachKeyPairResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachKeyPairResponseBodyResults) GetSuccess() *string {
	return s.Success
}

func (s *AttachKeyPairResponseBodyResults) SetCode(v string) *AttachKeyPairResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetInstanceId(v string) *AttachKeyPairResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetMessage(v string) *AttachKeyPairResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetSuccess(v string) *AttachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
