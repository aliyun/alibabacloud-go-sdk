// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataTrackResultDownloadStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *QueryDataTrackResultDownloadStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *QueryDataTrackResultDownloadStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryDataTrackResultDownloadStatusResponseBody
	GetRequestId() *string
	SetStatusResult(v *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) *QueryDataTrackResultDownloadStatusResponseBody
	GetStatusResult() *QueryDataTrackResultDownloadStatusResponseBodyStatusResult
	SetSuccess(v bool) *QueryDataTrackResultDownloadStatusResponseBody
	GetSuccess() *bool
}

type QueryDataTrackResultDownloadStatusResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the download progress.
	StatusResult *QueryDataTrackResultDownloadStatusResponseBodyStatusResult `json:"StatusResult,omitempty" xml:"StatusResult,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataTrackResultDownloadStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataTrackResultDownloadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) GetStatusResult() *QueryDataTrackResultDownloadStatusResponseBodyStatusResult {
	return s.StatusResult
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) SetErrorCode(v string) *QueryDataTrackResultDownloadStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) SetErrorMessage(v string) *QueryDataTrackResultDownloadStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) SetRequestId(v string) *QueryDataTrackResultDownloadStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) SetStatusResult(v *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) *QueryDataTrackResultDownloadStatusResponseBody {
	s.StatusResult = v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) SetSuccess(v bool) *QueryDataTrackResultDownloadStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDataTrackResultDownloadStatusResponseBodyStatusResult struct {
	// The status of the download task. Valid values:
	//
	// 	- **INIT**: The download task is being initialized.
	//
	// 	- **LISTING**: The download task is in a transient intermediate state during the initialization.
	//
	// 	- **DOWNLOADING**: The download task is being processed.
	//
	// 	- **DOWNLOAD_SUCCESS**: The download task was successfully processed.
	//
	// 	- **DOWNLOAD_FAIL**: The download task failed.
	//
	// example:
	//
	// DOWNLOAD_SUCCESS
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The URL that is used to download data tracking logs. This parameter is returned only when the value of DownloadStatus is DOWNLOAD_SUCCESS.
	//
	// example:
	//
	// https://idbsaasstore.oss-cn-zhangjiakou.aliyuncs.com/****_REDO_31201_207.zip?Expires=1682239593&OSSAccessKeyId=****&Signature=****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The description of the state.
	//
	// example:
	//
	// SUCCESS
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDataTrackResultDownloadStatusResponseBodyStatusResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDataTrackResultDownloadStatusResponseBodyStatusResult) GoString() string {
	return s.String()
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) GetDownloadStatus() *string {
	return s.DownloadStatus
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) SetDownloadStatus(v string) *QueryDataTrackResultDownloadStatusResponseBodyStatusResult {
	s.DownloadStatus = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) SetDownloadUrl(v string) *QueryDataTrackResultDownloadStatusResponseBodyStatusResult {
	s.DownloadUrl = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) SetStatusDesc(v string) *QueryDataTrackResultDownloadStatusResponseBodyStatusResult {
	s.StatusDesc = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) SetTotalCount(v int64) *QueryDataTrackResultDownloadStatusResponseBodyStatusResult {
	s.TotalCount = &v
	return s
}

func (s *QueryDataTrackResultDownloadStatusResponseBodyStatusResult) Validate() error {
	return dara.Validate(s)
}
