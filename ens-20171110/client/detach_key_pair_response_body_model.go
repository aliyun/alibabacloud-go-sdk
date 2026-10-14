// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailCount(v int32) *DetachKeyPairResponseBody
	GetFailCount() *int32
	SetKeyPairId(v string) *DetachKeyPairResponseBody
	GetKeyPairId() *string
	SetKeyPairName(v string) *DetachKeyPairResponseBody
	GetKeyPairName() *string
	SetRequestId(v string) *DetachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v []*DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody
	GetResults() []*DetachKeyPairResponseBodyResults
	SetTotalCount(v int32) *DetachKeyPairResponseBody
	GetTotalCount() *int32
}

type DetachKeyPairResponseBody struct {
	// The number of instances for which the detach operation failed.
	//
	// example:
	//
	// 1
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// ssh-xxx
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test-kp
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xx-xx-xx-xx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of the operation for each instance.
	Results []*DetachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of instances specified in the request.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBody) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DetachKeyPairResponseBody) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DetachKeyPairResponseBody) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachKeyPairResponseBody) GetResults() []*DetachKeyPairResponseBodyResults {
	return s.Results
}

func (s *DetachKeyPairResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DetachKeyPairResponseBody) SetFailCount(v int32) *DetachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetKeyPairId(v string) *DetachKeyPairResponseBody {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetKeyPairName(v string) *DetachKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetRequestId(v string) *DetachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetResults(v []*DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *DetachKeyPairResponseBody) SetTotalCount(v int32) *DetachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DetachKeyPairResponseBody) Validate() error {
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

type DetachKeyPairResponseBodyResults struct {
	// The status code for the operation on a specific instance. A value of `200` indicates a successful operation.
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
	// The response message for the operation. If the `Code` is `200`, the message is `successful`.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation on the instance succeeded. Valid values: `true` and `false`.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyResults) GetCode() *int32 {
	return s.Code
}

func (s *DetachKeyPairResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachKeyPairResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachKeyPairResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *DetachKeyPairResponseBodyResults) SetCode(v int32) *DetachKeyPairResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetInstanceId(v string) *DetachKeyPairResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetMessage(v string) *DetachKeyPairResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetSuccess(v bool) *DetachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
