// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEtlJobLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeEtlJobLogsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeEtlJobLogsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeEtlJobLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeEtlJobLogsResponseBody
	GetErrMessage() *string
	SetEtlRunningLogs(v []*DescribeEtlJobLogsResponseBodyEtlRunningLogs) *DescribeEtlJobLogsResponseBody
	GetEtlRunningLogs() []*DescribeEtlJobLogsResponseBodyEtlRunningLogs
	SetHttpStatusCode(v int32) *DescribeEtlJobLogsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeEtlJobLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEtlJobLogsResponseBody
	GetSuccess() *bool
}

type DescribeEtlJobLogsResponseBody struct {
	// The dynamic error code.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code. This example indicates that the specified ETL task ID is invalid.
	//
	// example:
	//
	// InvalidJobId
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message. This example indicates that the specified ETL task ID does not exist. In this case, the ETL task may be deleted.
	//
	// example:
	//
	// The specified dts job id %s is not exists.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The logs of ETL tasks.
	EtlRunningLogs []*DescribeEtlJobLogsResponseBodyEtlRunningLogs `json:"EtlRunningLogs,omitempty" xml:"EtlRunningLogs,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB43E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. If the call failed, false is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEtlJobLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEtlJobLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEtlJobLogsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeEtlJobLogsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeEtlJobLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeEtlJobLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeEtlJobLogsResponseBody) GetEtlRunningLogs() []*DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	return s.EtlRunningLogs
}

func (s *DescribeEtlJobLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeEtlJobLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEtlJobLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEtlJobLogsResponseBody) SetDynamicCode(v string) *DescribeEtlJobLogsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetDynamicMessage(v string) *DescribeEtlJobLogsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetErrCode(v string) *DescribeEtlJobLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetErrMessage(v string) *DescribeEtlJobLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetEtlRunningLogs(v []*DescribeEtlJobLogsResponseBodyEtlRunningLogs) *DescribeEtlJobLogsResponseBody {
	s.EtlRunningLogs = v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetHttpStatusCode(v int32) *DescribeEtlJobLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetRequestId(v string) *DescribeEtlJobLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) SetSuccess(v bool) *DescribeEtlJobLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBody) Validate() error {
	if s.EtlRunningLogs != nil {
		for _, item := range s.EtlRunningLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEtlJobLogsResponseBodyEtlRunningLogs struct {
	// The state of the ETL task.
	//
	// example:
	//
	// Starting DTS-ETL...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The module for which the logs are generated, such as the conversion module of ETL tasks.
	//
	// example:
	//
	// DTS-ETL
	ContentKey *string `json:"ContentKey,omitempty" xml:"ContentKey,omitempty"`
	// The ID of the ETL task.
	//
	// example:
	//
	// u**********5
	EtlId *string `json:"EtlId,omitempty" xml:"EtlId,omitempty"`
	// The time when the log was generated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1637306503000
	LogDatetime *string `json:"LogDatetime,omitempty" xml:"LogDatetime,omitempty"`
	// The log level. Valid values: ERROR, WARN, INFO, and DEBUG.
	//
	// example:
	//
	// INFO
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 121323*******454512
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEtlJobLogsResponseBodyEtlRunningLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeEtlJobLogsResponseBodyEtlRunningLogs) GoString() string {
	return s.String()
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetContent() *string {
	return s.Content
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetContentKey() *string {
	return s.ContentKey
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetEtlId() *string {
	return s.EtlId
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetLogDatetime() *string {
	return s.LogDatetime
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetStatus() *string {
	return s.Status
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetContent(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.Content = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetContentKey(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.ContentKey = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetEtlId(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.EtlId = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetLogDatetime(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.LogDatetime = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetStatus(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.Status = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) SetUserId(v string) *DescribeEtlJobLogsResponseBodyEtlRunningLogs {
	s.UserId = &v
	return s
}

func (s *DescribeEtlJobLogsResponseBodyEtlRunningLogs) Validate() error {
	return dara.Validate(s)
}
