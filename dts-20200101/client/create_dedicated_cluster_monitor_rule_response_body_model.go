// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedClusterMonitorRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateDedicatedClusterMonitorRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateDedicatedClusterMonitorRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *CreateDedicatedClusterMonitorRuleResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateDedicatedClusterMonitorRuleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDedicatedClusterMonitorRuleResponseBody
	GetSuccess() *string
}

type CreateDedicatedClusterMonitorRuleResponseBody struct {
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
	// The Value of Input Parameter %s is not valid.
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
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDedicatedClusterMonitorRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedClusterMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) SetErrCode(v string) *CreateDedicatedClusterMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) SetErrMessage(v string) *CreateDedicatedClusterMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) SetHttpStatusCode(v string) *CreateDedicatedClusterMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) SetRequestId(v string) *CreateDedicatedClusterMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) SetSuccess(v string) *CreateDedicatedClusterMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDedicatedClusterMonitorRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
