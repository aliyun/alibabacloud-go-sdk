// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetOverviewResponseBody
	GetCode() *int32
	SetData(v string) *GetOverviewResponseBody
	GetData() *string
	SetMessage(v string) *GetOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOverviewResponseBody
	GetSuccess() *bool
}

type GetOverviewResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The overview data in JSON format, which corresponds to the overview data on the console. The result is returned in one of the following three formats:
	//
	// - Basic information.
	//
	// - Node runtime information within a time interval.
	//
	// - Node runtime timing information within a time interval. This format returns statistics information at each time point for three data items: node triggers, successful executions, and failed executions.
	//
	// example:
	//
	// Basic info: {"schedulerx_job_counter_disable": "4","schedulerx_job_trigger_counter_running": "0","schedulerx_job_counter_enable": "70","schedulerx_job_counter_all": "74","schedulerx_worker_counter": "2"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// No access permission for the namespace [***]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOverviewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetOverviewResponseBody) GetData() *string {
	return s.Data
}

func (s *GetOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOverviewResponseBody) SetCode(v int32) *GetOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOverviewResponseBody) SetData(v string) *GetOverviewResponseBody {
	s.Data = &v
	return s
}

func (s *GetOverviewResponseBody) SetMessage(v string) *GetOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOverviewResponseBody) SetRequestId(v string) *GetOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOverviewResponseBody) SetSuccess(v bool) *GetOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}
