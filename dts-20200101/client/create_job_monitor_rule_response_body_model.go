// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobMonitorRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateJobMonitorRuleResponseBody
	GetCode() *string
	SetDtsJobId(v string) *CreateJobMonitorRuleResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *CreateJobMonitorRuleResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *CreateJobMonitorRuleResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateJobMonitorRuleResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateJobMonitorRuleResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateJobMonitorRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobMonitorRuleResponseBody
	GetSuccess() *bool
}

type CreateJobMonitorRuleResponseBody struct {
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic part in the error message. The value of this parameter is used to replace the \\*\\*%s\\*\\	- variable in the value of **ErrMessage**.
	//
	// > If the return value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the return value of **DynamicMessage*	- is **DtsJobId**, the specified value of **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 403
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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C498360-7892-433C-847A-BA71A850****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobMonitorRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateJobMonitorRuleResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateJobMonitorRuleResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateJobMonitorRuleResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateJobMonitorRuleResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateJobMonitorRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateJobMonitorRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobMonitorRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobMonitorRuleResponseBody) SetCode(v string) *CreateJobMonitorRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetDtsJobId(v string) *CreateJobMonitorRuleResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetDynamicMessage(v string) *CreateJobMonitorRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetErrCode(v string) *CreateJobMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetErrMessage(v string) *CreateJobMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetHttpStatusCode(v int32) *CreateJobMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetRequestId(v string) *CreateJobMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetSuccess(v bool) *CreateJobMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
