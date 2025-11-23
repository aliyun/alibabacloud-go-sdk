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
	// The download URL of the exported file.
	DownloadURLResult *GetDbExportDownloadURLResponseBodyDownloadURLResult `json:"DownloadURLResult,omitempty" xml:"DownloadURLResult,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request succeeded.
	//
	// 	- false: The request failed.
	//
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
	if s.DownloadURLResult != nil {
		if err := s.DownloadURLResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDbExportDownloadURLResponseBodyDownloadURLResult struct {
	// Indicates whether export results are available for download. Valid values:
	//
	// 	- **true**: Export results are available for download.
	//
	// 	- **false**: No export results are available for download.
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
	// The download URL of the exported file.
	//
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
