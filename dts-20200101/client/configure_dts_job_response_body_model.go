// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ConfigureDtsJobResponseBody
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ConfigureDtsJobResponseBody
	GetDtsJobId() *string
	SetErrCode(v string) *ConfigureDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ConfigureDtsJobResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ConfigureDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureDtsJobResponseBody
	GetSuccess() *string
}

type ConfigureDtsJobResponseBody struct {
	// The ID of the data migration or synchronization instance.
	//
	// example:
	//
	// dtsk2gm967v16f****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// k2gm967v16f****
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
	// The returned HTTP status codes.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobResponseBody) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ConfigureDtsJobResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConfigureDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureDtsJobResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ConfigureDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureDtsJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureDtsJobResponseBody) SetDtsInstanceId(v string) *ConfigureDtsJobResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetDtsJobId(v string) *ConfigureDtsJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetErrCode(v string) *ConfigureDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetErrMessage(v string) *ConfigureDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetHttpStatusCode(v string) *ConfigureDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetRequestId(v string) *ConfigureDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetSuccess(v string) *ConfigureDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
