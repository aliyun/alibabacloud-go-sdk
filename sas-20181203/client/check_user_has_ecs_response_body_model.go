// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserHasEcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckUserHasEcsResponseBody
	GetCode() *string
	SetData(v bool) *CheckUserHasEcsResponseBody
	GetData() *bool
	SetMessage(v string) *CheckUserHasEcsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckUserHasEcsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckUserHasEcsResponseBody
	GetSuccess() *bool
}

type CheckUserHasEcsResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request is successful. Other status codes indicate that the request fails. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether ECS instances exist. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserHasEcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUserHasEcsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserHasEcsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckUserHasEcsResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckUserHasEcsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckUserHasEcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUserHasEcsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckUserHasEcsResponseBody) SetCode(v string) *CheckUserHasEcsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckUserHasEcsResponseBody) SetData(v bool) *CheckUserHasEcsResponseBody {
	s.Data = &v
	return s
}

func (s *CheckUserHasEcsResponseBody) SetMessage(v string) *CheckUserHasEcsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckUserHasEcsResponseBody) SetRequestId(v string) *CheckUserHasEcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserHasEcsResponseBody) SetSuccess(v bool) *CheckUserHasEcsResponseBody {
	s.Success = &v
	return s
}

func (s *CheckUserHasEcsResponseBody) Validate() error {
	return dara.Validate(s)
}
