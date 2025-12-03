// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineScanReportUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineScanReportUrlResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineScanReportUrlResponseBody
	GetErrorMessage() *string
	SetReportUrl(v string) *GetPipelineScanReportUrlResponseBody
	GetReportUrl() *string
	SetRequestId(v string) *GetPipelineScanReportUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineScanReportUrlResponseBody
	GetSuccess() *bool
}

type GetPipelineScanReportUrlResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// http://aliyun.com
	ReportUrl *string `json:"reportUrl,omitempty" xml:"reportUrl,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineScanReportUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineScanReportUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineScanReportUrlResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineScanReportUrlResponseBody) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GetPipelineScanReportUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineScanReportUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineScanReportUrlResponseBody) SetErrorCode(v string) *GetPipelineScanReportUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetErrorMessage(v string) *GetPipelineScanReportUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetReportUrl(v string) *GetPipelineScanReportUrlResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetRequestId(v string) *GetPipelineScanReportUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) SetSuccess(v bool) *GetPipelineScanReportUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineScanReportUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
