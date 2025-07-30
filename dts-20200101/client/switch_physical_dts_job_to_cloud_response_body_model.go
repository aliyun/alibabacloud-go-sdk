// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchPhysicalDtsJobToCloudResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *SwitchPhysicalDtsJobToCloudResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SwitchPhysicalDtsJobToCloudResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *SwitchPhysicalDtsJobToCloudResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SwitchPhysicalDtsJobToCloudResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SwitchPhysicalDtsJobToCloudResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SwitchPhysicalDtsJobToCloudResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchPhysicalDtsJobToCloudResponseBody
	GetSuccess() *bool
}

type SwitchPhysicalDtsJobToCloudResponseBody struct {
	// Dynamic error code, this parameter will be deprecated.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message used to replace the **%s*	- in the **ErrMessage*	- return parameter.  > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Returns the corresponding error message when an invocation error occurs.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// request ID.
	//
	// example:
	//
	// 659304E3-D44E-5EFA-BDE3-60015E30403B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchPhysicalDtsJobToCloudResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchPhysicalDtsJobToCloudResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetDynamicCode(v string) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetDynamicMessage(v string) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetErrCode(v string) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetErrMessage(v string) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetHttpStatusCode(v int32) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetRequestId(v string) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) SetSuccess(v bool) *SwitchPhysicalDtsJobToCloudResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchPhysicalDtsJobToCloudResponseBody) Validate() error {
	return dara.Validate(s)
}
