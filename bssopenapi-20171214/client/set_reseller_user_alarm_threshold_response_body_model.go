// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserAlarmThresholdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetResellerUserAlarmThresholdResponseBody
	GetCode() *string
	SetData(v bool) *SetResellerUserAlarmThresholdResponseBody
	GetData() *bool
	SetMessage(v string) *SetResellerUserAlarmThresholdResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetResellerUserAlarmThresholdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetResellerUserAlarmThresholdResponseBody
	GetSuccess() *bool
}

type SetResellerUserAlarmThresholdResponseBody struct {
	// The error code returned when the call fails. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates that the call was completed. A value of false indicates that an error occurred.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetResellerUserAlarmThresholdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserAlarmThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetResellerUserAlarmThresholdResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetResellerUserAlarmThresholdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetResellerUserAlarmThresholdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetResellerUserAlarmThresholdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetCode(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetData(v bool) *SetResellerUserAlarmThresholdResponseBody {
	s.Data = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetMessage(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetRequestId(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetSuccess(v bool) *SetResellerUserAlarmThresholdResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) Validate() error {
	return dara.Validate(s)
}
