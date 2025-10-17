// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsServiceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeDtsServiceLogResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDtsServiceLogResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDtsServiceLogResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDtsServiceLogResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeDtsServiceLogResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeDtsServiceLogResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDtsServiceLogResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDtsServiceLogResponseBody
	GetRequestId() *string
	SetServiceLogContexts(v []*DescribeDtsServiceLogResponseBodyServiceLogContexts) *DescribeDtsServiceLogResponseBody
	GetServiceLogContexts() []*DescribeDtsServiceLogResponseBodyServiceLogContexts
	SetSuccess(v bool) *DescribeDtsServiceLogResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int64) *DescribeDtsServiceLogResponseBody
	GetTotalRecordCount() *int64
}

type DescribeDtsServiceLogResponseBody struct {
	// The dynamic error code. This parameter will be removed soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries returned per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F87DF250-952C-47FE-8A02-69414FAA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the logs.
	ServiceLogContexts []*DescribeDtsServiceLogResponseBodyServiceLogContexts `json:"ServiceLogContexts,omitempty" xml:"ServiceLogContexts,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of logs that meet the query conditions.
	//
	// example:
	//
	// 35
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDtsServiceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsServiceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsServiceLogResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDtsServiceLogResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDtsServiceLogResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDtsServiceLogResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsServiceLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDtsServiceLogResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsServiceLogResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDtsServiceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDtsServiceLogResponseBody) GetServiceLogContexts() []*DescribeDtsServiceLogResponseBodyServiceLogContexts {
	return s.ServiceLogContexts
}

func (s *DescribeDtsServiceLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDtsServiceLogResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeDtsServiceLogResponseBody) SetDynamicCode(v string) *DescribeDtsServiceLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetDynamicMessage(v string) *DescribeDtsServiceLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetErrCode(v string) *DescribeDtsServiceLogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetErrMessage(v string) *DescribeDtsServiceLogResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetHttpStatusCode(v int32) *DescribeDtsServiceLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetPageNumber(v int32) *DescribeDtsServiceLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetPageRecordCount(v int32) *DescribeDtsServiceLogResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetRequestId(v string) *DescribeDtsServiceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetServiceLogContexts(v []*DescribeDtsServiceLogResponseBodyServiceLogContexts) *DescribeDtsServiceLogResponseBody {
	s.ServiceLogContexts = v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetSuccess(v bool) *DescribeDtsServiceLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) SetTotalRecordCount(v int64) *DescribeDtsServiceLogResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBody) Validate() error {
	if s.ServiceLogContexts != nil {
		for _, item := range s.ServiceLogContexts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDtsServiceLogResponseBodyServiceLogContexts struct {
	// The log content.
	//
	// example:
	//
	// Statistics: generator = 369173; collector = 470109; replicator = 2470; ping = 2/2/2; execute = 29/29/29; rt = 29/29/29; state = IDLE; queries = -1; exceptions = {connects = 0, replicates = 0}; infos = {}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The log level.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the logs were collected. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-13T09:13:39.443+00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDtsServiceLogResponseBodyServiceLogContexts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsServiceLogResponseBodyServiceLogContexts) GoString() string {
	return s.String()
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) GetContext() *string {
	return s.Context
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) GetState() *string {
	return s.State
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) GetTime() *string {
	return s.Time
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) SetContext(v string) *DescribeDtsServiceLogResponseBodyServiceLogContexts {
	s.Context = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) SetState(v string) *DescribeDtsServiceLogResponseBodyServiceLogContexts {
	s.State = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) SetTime(v string) *DescribeDtsServiceLogResponseBodyServiceLogContexts {
	s.Time = &v
	return s
}

func (s *DescribeDtsServiceLogResponseBodyServiceLogContexts) Validate() error {
	return dara.Validate(s)
}
