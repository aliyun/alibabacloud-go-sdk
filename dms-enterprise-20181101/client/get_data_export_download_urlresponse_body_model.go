// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportDownloadURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadURLResult(v *GetDataExportDownloadURLResponseBodyDownloadURLResult) *GetDataExportDownloadURLResponseBody
	GetDownloadURLResult() *GetDataExportDownloadURLResponseBodyDownloadURLResult
	SetErrorCode(v string) *GetDataExportDownloadURLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataExportDownloadURLResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataExportDownloadURLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataExportDownloadURLResponseBody
	GetSuccess() *bool
}

type GetDataExportDownloadURLResponseBody struct {
	// The details of the download URL of the file that records the export results for the ticket.
	DownloadURLResult *GetDataExportDownloadURLResponseBodyDownloadURLResult `json:"DownloadURLResult,omitempty" xml:"DownloadURLResult,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4BF24EA5-9013-4C85-AE68-6C23AF5E0097
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

func (s GetDataExportDownloadURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportDownloadURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponseBody) GetDownloadURLResult() *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	return s.DownloadURLResult
}

func (s *GetDataExportDownloadURLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataExportDownloadURLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataExportDownloadURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataExportDownloadURLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataExportDownloadURLResponseBody) SetDownloadURLResult(v *GetDataExportDownloadURLResponseBodyDownloadURLResult) *GetDataExportDownloadURLResponseBody {
	s.DownloadURLResult = v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetErrorCode(v string) *GetDataExportDownloadURLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetErrorMessage(v string) *GetDataExportDownloadURLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetRequestId(v string) *GetDataExportDownloadURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetSuccess(v bool) *GetDataExportDownloadURLResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataExportDownloadURLResponseBodyDownloadURLResult struct {
	// Indicates whether export results are available for download. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasResult *bool `json:"HasResult,omitempty" xml:"HasResult,omitempty"`
	// The message that indicates an exception.
	//
	// example:
	//
	// tip message
	TipMessage *string `json:"TipMessage,omitempty" xml:"TipMessage,omitempty"`
	// The download URL of the file that records the export results for the ticket.
	//
	// example:
	//
	// https://dms-idb-hangzhou.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetDataExportDownloadURLResponseBodyDownloadURLResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportDownloadURLResponseBodyDownloadURLResult) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) GetHasResult() *bool {
	return s.HasResult
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) GetTipMessage() *string {
	return s.TipMessage
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) GetURL() *string {
	return s.URL
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetHasResult(v bool) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.HasResult = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetTipMessage(v string) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.TipMessage = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetURL(v string) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.URL = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) Validate() error {
	return dara.Validate(s)
}
