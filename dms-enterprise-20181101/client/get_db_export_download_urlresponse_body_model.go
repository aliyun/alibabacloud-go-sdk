// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbExportDownloadURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadURLResult(v *GetDbExportDownloadURLResponseBodyDownloadURLResult) *GetDbExportDownloadURLResponseBody
	GetDownloadURLResult() *GetDbExportDownloadURLResponseBodyDownloadURLResult
	SetErrorCode(v string) *GetDbExportDownloadURLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDbExportDownloadURLResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDbExportDownloadURLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDbExportDownloadURLResponseBody
	GetSuccess() *bool
}

type GetDbExportDownloadURLResponseBody struct {
	DownloadURLResult *GetDbExportDownloadURLResponseBodyDownloadURLResult `json:"DownloadURLResult,omitempty" xml:"DownloadURLResult,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDbExportDownloadURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDbExportDownloadURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetDbExportDownloadURLResponseBody) GetDownloadURLResult() *GetDbExportDownloadURLResponseBodyDownloadURLResult {
	return s.DownloadURLResult
}

func (s *GetDbExportDownloadURLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDbExportDownloadURLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDbExportDownloadURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDbExportDownloadURLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDbExportDownloadURLResponseBody) SetDownloadURLResult(v *GetDbExportDownloadURLResponseBodyDownloadURLResult) *GetDbExportDownloadURLResponseBody {
	s.DownloadURLResult = v
	return s
}

func (s *GetDbExportDownloadURLResponseBody) SetErrorCode(v string) *GetDbExportDownloadURLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBody) SetErrorMessage(v string) *GetDbExportDownloadURLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBody) SetRequestId(v string) *GetDbExportDownloadURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBody) SetSuccess(v bool) *GetDbExportDownloadURLResponseBody {
	s.Success = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDbExportDownloadURLResponseBodyDownloadURLResult struct {
	// example:
	//
	// true
	HasResult *bool `json:"HasResult,omitempty" xml:"HasResult,omitempty"`
	// example:
	//
	// tip message
	TipMessage *string `json:"TipMessage,omitempty" xml:"TipMessage,omitempty"`
	// example:
	//
	// https://dms-idb-hangzhou.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetDbExportDownloadURLResponseBodyDownloadURLResult) String() string {
	return dara.Prettify(s)
}

func (s GetDbExportDownloadURLResponseBodyDownloadURLResult) GoString() string {
	return s.String()
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) GetHasResult() *bool {
	return s.HasResult
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) GetTipMessage() *string {
	return s.TipMessage
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) GetURL() *string {
	return s.URL
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) SetHasResult(v bool) *GetDbExportDownloadURLResponseBodyDownloadURLResult {
	s.HasResult = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) SetTipMessage(v string) *GetDbExportDownloadURLResponseBodyDownloadURLResult {
	s.TipMessage = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) SetURL(v string) *GetDbExportDownloadURLResponseBodyDownloadURLResult {
	s.URL = &v
	return s
}

func (s *GetDbExportDownloadURLResponseBodyDownloadURLResult) Validate() error {
	return dara.Validate(s)
}
