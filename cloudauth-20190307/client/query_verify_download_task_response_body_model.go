// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *QueryVerifyDownloadTaskResponseBody
	GetErrorCode() *string
	SetFinish(v bool) *QueryVerifyDownloadTaskResponseBody
	GetFinish() *bool
	SetRequestId(v string) *QueryVerifyDownloadTaskResponseBody
	GetRequestId() *string
	SetStatus(v int32) *QueryVerifyDownloadTaskResponseBody
	GetStatus() *int32
	SetUrl(v string) *QueryVerifyDownloadTaskResponseBody
	GetUrl() *string
}

type QueryVerifyDownloadTaskResponseBody struct {
	// Error code.
	//
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Whether the download task is completed:
	//
	// - **true**: Completed
	//
	// - **false**: Not completed
	//
	// example:
	//
	// true
	Finish *bool `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task status:
	//
	// - **1**: File generation in progress
	//
	// - **2**: File generation completed
	//
	// - **3**: File generation failed
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Download URL.
	//
	// example:
	//
	// http://xxx/xxx.csv
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryVerifyDownloadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVerifyDownloadTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryVerifyDownloadTaskResponseBody) GetFinish() *bool {
	return s.Finish
}

func (s *QueryVerifyDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVerifyDownloadTaskResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *QueryVerifyDownloadTaskResponseBody) GetUrl() *string {
	return s.Url
}

func (s *QueryVerifyDownloadTaskResponseBody) SetErrorCode(v string) *QueryVerifyDownloadTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponseBody) SetFinish(v bool) *QueryVerifyDownloadTaskResponseBody {
	s.Finish = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponseBody) SetRequestId(v string) *QueryVerifyDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponseBody) SetStatus(v int32) *QueryVerifyDownloadTaskResponseBody {
	s.Status = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponseBody) SetUrl(v string) *QueryVerifyDownloadTaskResponseBody {
	s.Url = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
