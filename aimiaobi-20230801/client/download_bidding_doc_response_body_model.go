// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadBiddingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadBiddingDocResponseBody
	GetCode() *string
	SetData(v *DownloadBiddingDocResponseBodyData) *DownloadBiddingDocResponseBody
	GetData() *DownloadBiddingDocResponseBodyData
	SetHttpStatusCode(v int32) *DownloadBiddingDocResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DownloadBiddingDocResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadBiddingDocResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadBiddingDocResponseBody
	GetSuccess() *bool
}

type DownloadBiddingDocResponseBody struct {
	// Status code
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *DownloadBiddingDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Operation result: true for success, false for failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadBiddingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadBiddingDocResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadBiddingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadBiddingDocResponseBody) GetData() *DownloadBiddingDocResponseBodyData {
	return s.Data
}

func (s *DownloadBiddingDocResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DownloadBiddingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadBiddingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadBiddingDocResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadBiddingDocResponseBody) SetCode(v string) *DownloadBiddingDocResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadBiddingDocResponseBody) SetData(v *DownloadBiddingDocResponseBodyData) *DownloadBiddingDocResponseBody {
	s.Data = v
	return s
}

func (s *DownloadBiddingDocResponseBody) SetHttpStatusCode(v int32) *DownloadBiddingDocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadBiddingDocResponseBody) SetMessage(v string) *DownloadBiddingDocResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadBiddingDocResponseBody) SetRequestId(v string) *DownloadBiddingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadBiddingDocResponseBody) SetSuccess(v bool) *DownloadBiddingDocResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadBiddingDocResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadBiddingDocResponseBodyData struct {
	// Task ID
	//
	// example:
	//
	// 111_Default_20250708142918
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Document URL
	//
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadBiddingDocResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DownloadBiddingDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *DownloadBiddingDocResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DownloadBiddingDocResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *DownloadBiddingDocResponseBodyData) SetTaskId(v string) *DownloadBiddingDocResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DownloadBiddingDocResponseBodyData) SetUrl(v string) *DownloadBiddingDocResponseBodyData {
	s.Url = &v
	return s
}

func (s *DownloadBiddingDocResponseBodyData) Validate() error {
	return dara.Validate(s)
}
