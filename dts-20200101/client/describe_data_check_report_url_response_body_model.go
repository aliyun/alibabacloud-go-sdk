// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckReportUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicMessage(v string) *DescribeDataCheckReportUrlResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDataCheckReportUrlResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDataCheckReportUrlResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeDataCheckReportUrlResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeDataCheckReportUrlResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDataCheckReportUrlResponseBody
	GetSuccess() *string
}

type DescribeDataCheckReportUrlResponseBody struct {
	// The URL for downloading the verification report.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF2DE23-B692-5B85-90B7-44B6F4D8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDataCheckReportUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckReportUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckReportUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDataCheckReportUrlResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDataCheckReportUrlResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDataCheckReportUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDataCheckReportUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataCheckReportUrlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDataCheckReportUrlResponseBody) SetDynamicMessage(v string) *DescribeDataCheckReportUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) SetErrCode(v string) *DescribeDataCheckReportUrlResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) SetErrMessage(v string) *DescribeDataCheckReportUrlResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) SetHttpStatusCode(v int32) *DescribeDataCheckReportUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) SetRequestId(v string) *DescribeDataCheckReportUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) SetSuccess(v string) *DescribeDataCheckReportUrlResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
