// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ConfigureSubscriptionResponseBody
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ConfigureSubscriptionResponseBody
	GetDtsJobId() *string
	SetErrCode(v string) *ConfigureSubscriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureSubscriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ConfigureSubscriptionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ConfigureSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureSubscriptionResponseBody
	GetSuccess() *string
}

type ConfigureSubscriptionResponseBody struct {
	// The ID of the change tracking instance.
	//
	// example:
	//
	// dtsy0zz3t13h7d****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task.
	//
	// example:
	//
	// y0zz3t13h7d****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1D6ECADF-C5E9-4C96-8811-77602B31****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionResponseBody) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ConfigureSubscriptionResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConfigureSubscriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureSubscriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureSubscriptionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ConfigureSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureSubscriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureSubscriptionResponseBody) SetDtsInstanceId(v string) *ConfigureSubscriptionResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetDtsJobId(v string) *ConfigureSubscriptionResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetErrCode(v string) *ConfigureSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetErrMessage(v string) *ConfigureSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetHttpStatusCode(v string) *ConfigureSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetRequestId(v string) *ConfigureSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetSuccess(v string) *ConfigureSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
