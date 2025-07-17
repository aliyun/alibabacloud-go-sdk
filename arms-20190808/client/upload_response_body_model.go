// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadResponseBody
	GetRequestId() *string
	SetUploadResult(v *UploadResponseBodyUploadResult) *UploadResponseBody
	GetUploadResult() *UploadResponseBodyUploadResult
}

type UploadResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned data.
	UploadResult *UploadResponseBodyUploadResult `json:"UploadResult,omitempty" xml:"UploadResult,omitempty" type:"Struct"`
}

func (s UploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadResponseBody) GoString() string {
	return s.String()
}

func (s *UploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadResponseBody) GetUploadResult() *UploadResponseBodyUploadResult {
	return s.UploadResult
}

func (s *UploadResponseBody) SetRequestId(v string) *UploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadResponseBody) SetUploadResult(v *UploadResponseBodyUploadResult) *UploadResponseBody {
	s.UploadResult = v
	return s
}

func (s *UploadResponseBody) Validate() error {
	return dara.Validate(s)
}

type UploadResponseBodyUploadResult struct {
	// The ID of the SourceMap file.
	//
	// example:
	//
	// 123
	Fid *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	// The name of the SourceMap file.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the file was uploaded.
	//
	// example:
	//
	// 1650272251
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s UploadResponseBodyUploadResult) String() string {
	return dara.Prettify(s)
}

func (s UploadResponseBodyUploadResult) GoString() string {
	return s.String()
}

func (s *UploadResponseBodyUploadResult) GetFid() *string {
	return s.Fid
}

func (s *UploadResponseBodyUploadResult) GetFileName() *string {
	return s.FileName
}

func (s *UploadResponseBodyUploadResult) GetUploadTime() *string {
	return s.UploadTime
}

func (s *UploadResponseBodyUploadResult) SetFid(v string) *UploadResponseBodyUploadResult {
	s.Fid = &v
	return s
}

func (s *UploadResponseBodyUploadResult) SetFileName(v string) *UploadResponseBodyUploadResult {
	s.FileName = &v
	return s
}

func (s *UploadResponseBodyUploadResult) SetUploadTime(v string) *UploadResponseBodyUploadResult {
	s.UploadTime = &v
	return s
}

func (s *UploadResponseBodyUploadResult) Validate() error {
	return dara.Validate(s)
}
