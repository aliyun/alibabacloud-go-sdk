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
	SetRequestId(v string) *DetachKeyPairResponseBody
	GetRequestId() *string
	SetResults(v []*DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody
	GetResults() []*DetachKeyPairResponseBodyResults
	SetTotalCount(v int32) *DetachKeyPairResponseBody
	GetTotalCount() *int32
}

type DetachKeyPairResponseBody struct {
	// The total number of simple application servers from which you fail to unbind key pairs.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request results.
	Results []*DetachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of simple application servers from which the SSH key pair is unbound.
	//
	// example:
	//
	// 1
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// aa6e71ddb35c46679bc4753d6219d604
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the key pair is unbound from the simple application server successfully. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachKeyPairResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachKeyPairResponseBodyResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachKeyPairResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachKeyPairResponseBodyResults) GetSuccess() *string {
	return s.Success
}

func (s *DetachKeyPairResponseBodyResults) SetCode(v string) *DetachKeyPairResponseBodyResults {
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

func (s *DetachKeyPairResponseBodyResults) SetSuccess(v string) *DetachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
