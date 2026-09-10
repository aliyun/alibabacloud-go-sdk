// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *GetDataCheckReportStatusResponseBody
	GetData() *int32
	SetErrCode(v string) *GetDataCheckReportStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckReportStatusResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckReportStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckReportStatusResponseBody
	GetSuccess() *bool
}

type GetDataCheckReportStatusResponseBody struct {
	// The report generation status. Valid values:
	//
	// - 0: not generated
	//
	// - 1: generating
	//
	// - 2: generated
	//
	// - 3: generation failed
	//
	// example:
	//
	// 0
	Data *int32 `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckReportStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportStatusResponseBody) GetData() *int32 {
	return s.Data
}

func (s *GetDataCheckReportStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckReportStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckReportStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckReportStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckReportStatusResponseBody) SetData(v int32) *GetDataCheckReportStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetDataCheckReportStatusResponseBody) SetErrCode(v string) *GetDataCheckReportStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckReportStatusResponseBody) SetErrMessage(v string) *GetDataCheckReportStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckReportStatusResponseBody) SetRequestId(v string) *GetDataCheckReportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckReportStatusResponseBody) SetSuccess(v bool) *GetDataCheckReportStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckReportStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
