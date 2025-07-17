// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumUploadFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetRumUploadFilesResponseBody
	GetCode() *int32
	SetData(v []*GetRumUploadFilesResponseBodyData) *GetRumUploadFilesResponseBody
	GetData() []*GetRumUploadFilesResponseBodyData
	SetHttpStatusCode(v int32) *GetRumUploadFilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRumUploadFilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumUploadFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRumUploadFilesResponseBody
	GetSuccess() *bool
}

type GetRumUploadFilesResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried files.
	Data []*GetRumUploadFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2983BEF7-4A0D-47A2-94A2-8E9C5E63****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRumUploadFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumUploadFilesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumUploadFilesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetRumUploadFilesResponseBody) GetData() []*GetRumUploadFilesResponseBodyData {
	return s.Data
}

func (s *GetRumUploadFilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRumUploadFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumUploadFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumUploadFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRumUploadFilesResponseBody) SetCode(v int32) *GetRumUploadFilesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumUploadFilesResponseBody) SetData(v []*GetRumUploadFilesResponseBodyData) *GetRumUploadFilesResponseBody {
	s.Data = v
	return s
}

func (s *GetRumUploadFilesResponseBody) SetHttpStatusCode(v int32) *GetRumUploadFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRumUploadFilesResponseBody) SetMessage(v string) *GetRumUploadFilesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumUploadFilesResponseBody) SetRequestId(v string) *GetRumUploadFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumUploadFilesResponseBody) SetSuccess(v bool) *GetRumUploadFilesResponseBody {
	s.Success = &v
	return s
}

func (s *GetRumUploadFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRumUploadFilesResponseBodyData struct {
	// The file name.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the file was last modified. The value is a timestamp.
	//
	// example:
	//
	// 1682863151000
	LastModifiedTime interface{} `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 1000
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The file ID.
	//
	// example:
	//
	// cxxxerfewrfwerfwerffvrt
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version number of the file.
	//
	// example:
	//
	// 1.0.0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetRumUploadFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRumUploadFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRumUploadFilesResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetRumUploadFilesResponseBodyData) GetLastModifiedTime() interface{} {
	return s.LastModifiedTime
}

func (s *GetRumUploadFilesResponseBodyData) GetSize() *string {
	return s.Size
}

func (s *GetRumUploadFilesResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetRumUploadFilesResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *GetRumUploadFilesResponseBodyData) SetFileName(v string) *GetRumUploadFilesResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetRumUploadFilesResponseBodyData) SetLastModifiedTime(v interface{}) *GetRumUploadFilesResponseBodyData {
	s.LastModifiedTime = v
	return s
}

func (s *GetRumUploadFilesResponseBodyData) SetSize(v string) *GetRumUploadFilesResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetRumUploadFilesResponseBodyData) SetUuid(v string) *GetRumUploadFilesResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetRumUploadFilesResponseBodyData) SetVersionId(v string) *GetRumUploadFilesResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *GetRumUploadFilesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
