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
	SetKeyPairId(v string) *AttachKeyPairResponseBody
	GetKeyPairId() *string
	SetKeyPairName(v string) *AttachKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *AttachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v []*AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody
	GetResults() []*AttachKeyPairResponseBodyResults
	SetTotalCount(v int32) *AttachKeyPairResponseBody
	GetTotalCount() *int32
}

type AttachKeyPairResponseBody struct {
	// The number of instances from which the SSH key pair failed to be unbound.
	//
	// example:
	//
	// 1
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the SSH key pair.
	//
	// example:
	//
	// ssh-xxx
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the SSH key pair.
	//
	// example:
	//
	// ssh-xxx
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xx-xx-xx-xx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result set of the unbind operation.
	Results []*AttachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of instances from which the SSH key pair is unbound.
	//
	// example:
	//
	// 3
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

func (s *AttachKeyPairResponseBody) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *AttachKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
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

func (s *AttachKeyPairResponseBody) SetKeyPairId(v string) *AttachKeyPairResponseBody {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetKeyPairName(v string) *AttachKeyPairResponseBody {
	s.KeyPairName = &v
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
	// The status code of the operation. 200 indicates that the operation succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The message of the operation. If the value of the Code parameter is 200, the value of this parameter is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyResults) GetCode() *int32 {
	return s.Code
}

func (s *AttachKeyPairResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachKeyPairResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachKeyPairResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *AttachKeyPairResponseBodyResults) SetCode(v int32) *AttachKeyPairResponseBodyResults {
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

func (s *AttachKeyPairResponseBodyResults) SetSuccess(v bool) *AttachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
