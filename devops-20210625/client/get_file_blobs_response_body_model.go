// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileBlobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetFileBlobsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFileBlobsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetFileBlobsResponseBody
	GetRequestId() *string
	SetResult(v *GetFileBlobsResponseBodyResult) *GetFileBlobsResponseBody
	GetResult() *GetFileBlobsResponseBodyResult
	SetSuccess(v bool) *GetFileBlobsResponseBody
	GetSuccess() *bool
}

type GetFileBlobsResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F590C9D8-E908-5B6C-95AC-56B7E8011FFA
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetFileBlobsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFileBlobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileBlobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFileBlobsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFileBlobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileBlobsResponseBody) GetResult() *GetFileBlobsResponseBodyResult {
	return s.Result
}

func (s *GetFileBlobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileBlobsResponseBody) SetErrorCode(v string) *GetFileBlobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetErrorMessage(v string) *GetFileBlobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetRequestId(v string) *GetFileBlobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetResult(v *GetFileBlobsResponseBodyResult) *GetFileBlobsResponseBody {
	s.Result = v
	return s
}

func (s *GetFileBlobsResponseBody) SetSuccess(v bool) *GetFileBlobsResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileBlobsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileBlobsResponseBodyResult struct {
	// example:
	//
	// xxxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Size    *int64  `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 65535
	TotalLines *int32 `json:"totalLines,omitempty" xml:"totalLines,omitempty"`
}

func (s GetFileBlobsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFileBlobsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetFileBlobsResponseBodyResult) GetSize() *int64 {
	return s.Size
}

func (s *GetFileBlobsResponseBodyResult) GetTotalLines() *int32 {
	return s.TotalLines
}

func (s *GetFileBlobsResponseBodyResult) SetContent(v string) *GetFileBlobsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) SetSize(v int64) *GetFileBlobsResponseBodyResult {
	s.Size = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) SetTotalLines(v int32) *GetFileBlobsResponseBodyResult {
	s.TotalLines = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
