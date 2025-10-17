// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterOperateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeClusterOperateLogsResponseBody
	GetCode() *string
	SetDataPoints(v []*DescribeClusterOperateLogsResponseBodyDataPoints) *DescribeClusterOperateLogsResponseBody
	GetDataPoints() []*DescribeClusterOperateLogsResponseBodyDataPoints
	SetDynamicMessage(v string) *DescribeClusterOperateLogsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeClusterOperateLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeClusterOperateLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeClusterOperateLogsResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeClusterOperateLogsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeClusterOperateLogsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeClusterOperateLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeClusterOperateLogsResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int64) *DescribeClusterOperateLogsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeClusterOperateLogsResponseBody struct {
	// The error code returned by the backend service. The number is incremented.
	//
	// example:
	//
	// 500
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the operation logs.
	DataPoints []*DescribeClusterOperateLogsResponseBodyDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace %s in the ErrMessage parameter.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned for an exception.
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
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeClusterOperateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperateLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeClusterOperateLogsResponseBody) GetDataPoints() []*DescribeClusterOperateLogsResponseBodyDataPoints {
	return s.DataPoints
}

func (s *DescribeClusterOperateLogsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeClusterOperateLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeClusterOperateLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeClusterOperateLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeClusterOperateLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterOperateLogsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeClusterOperateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterOperateLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeClusterOperateLogsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeClusterOperateLogsResponseBody) SetCode(v string) *DescribeClusterOperateLogsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetDataPoints(v []*DescribeClusterOperateLogsResponseBodyDataPoints) *DescribeClusterOperateLogsResponseBody {
	s.DataPoints = v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetDynamicMessage(v string) *DescribeClusterOperateLogsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetErrCode(v string) *DescribeClusterOperateLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetErrMessage(v string) *DescribeClusterOperateLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetHttpStatusCode(v int32) *DescribeClusterOperateLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetPageNumber(v int32) *DescribeClusterOperateLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetPageRecordCount(v int32) *DescribeClusterOperateLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetRequestId(v string) *DescribeClusterOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetSuccess(v bool) *DescribeClusterOperateLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) SetTotalRecordCount(v int64) *DescribeClusterOperateLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBody) Validate() error {
	if s.DataPoints != nil {
		for _, item := range s.DataPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterOperateLogsResponseBodyDataPoints struct {
	// Other description of the operation.
	//
	// example:
	//
	// null
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The primary key of the log table.
	//
	// example:
	//
	// 237827
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the operation is recorded. The value is in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1642077211574
	LogDatetime *int64 `json:"LogDatetime,omitempty" xml:"LogDatetime,omitempty"`
	// The value of the parameter after the operation if the operation is an UPDATE operation.
	//
	// example:
	//
	// 105
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The value of the parameter before the operation if the operation is an UPDATE operation.
	//
	// example:
	//
	// 100
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	// The type of the operation.
	//
	// example:
	//
	// modify-oversold-ratio
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The user who performed the operation.
	//
	// example:
	//
	// null
	OperationUser *string `json:"OperationUser,omitempty" xml:"OperationUser,omitempty"`
	// Indicates whether the request was successful. A value of **1*	- indicates that the request was successful.
	//
	// example:
	//
	// 1
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeClusterOperateLogsResponseBodyDataPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterOperateLogsResponseBodyDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetContent() *string {
	return s.Content
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetId() *string {
	return s.Id
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetLogDatetime() *int64 {
	return s.LogDatetime
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetOldValue() *string {
	return s.OldValue
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetOperationName() *string {
	return s.OperationName
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetOperationUser() *string {
	return s.OperationUser
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) GetSuccess() *int32 {
	return s.Success
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetContent(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.Content = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetId(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.Id = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetLogDatetime(v int64) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.LogDatetime = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetNewValue(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.NewValue = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetOldValue(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.OldValue = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetOperationName(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.OperationName = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetOperationUser(v string) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.OperationUser = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) SetSuccess(v int32) *DescribeClusterOperateLogsResponseBodyDataPoints {
	s.Success = &v
	return s
}

func (s *DescribeClusterOperateLogsResponseBodyDataPoints) Validate() error {
	return dara.Validate(s)
}
