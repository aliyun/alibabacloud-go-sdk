// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateUploadUrlResponseBody
	GetCode() *string
	SetData(v *GenerateUploadUrlResponseBodyData) *GenerateUploadUrlResponseBody
	GetData() *GenerateUploadUrlResponseBodyData
	SetHttpStatusCode(v int32) *GenerateUploadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateUploadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateUploadUrlResponseBody
	GetSuccess() *bool
}

type GenerateUploadUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateUploadUrlResponseBody) GetData() *GenerateUploadUrlResponseBodyData {
	return s.Data
}

func (s *GenerateUploadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateUploadUrlResponseBody) SetCode(v string) *GenerateUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetData(v *GenerateUploadUrlResponseBodyData) *GenerateUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetHttpStatusCode(v int32) *GenerateUploadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetMessage(v string) *GenerateUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetRequestId(v string) *GenerateUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) SetSuccess(v bool) *GenerateUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateUploadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateUploadUrlResponseBodyData struct {
	// example:
	//
	// LTAIvKWEr4DoFSqz
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 1742960933
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// pload/file/874d7371-791b-4281-935c-637630a37785/874d7371-791b-4281-935c-637630a37785_9bd7916d-a340-4925-a911-92390cbe0f33.json
	Folder *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	// example:
	//
	// https://cloudagentbot-online.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0yNlQwMzo0ODo1My4wMzVaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNTI0Mjg4MF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJ1cGxvYWQvZmlsZS84NzRkNzM3MS03OTFiLTQyODEtOTM1Yy02Mzc2MzBhMzc3ODUvIl1dfQ==",
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// I6d1ONWVuTj5i0Kz4Vn+V0lC6v4=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *GenerateUploadUrlResponseBodyData) GetExpire() *int32 {
	return s.Expire
}

func (s *GenerateUploadUrlResponseBodyData) GetFolder() *string {
	return s.Folder
}

func (s *GenerateUploadUrlResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GenerateUploadUrlResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GenerateUploadUrlResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GenerateUploadUrlResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GenerateUploadUrlResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateUploadUrlResponseBodyData) SetAccessId(v string) *GenerateUploadUrlResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetExpire(v int32) *GenerateUploadUrlResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetFolder(v string) *GenerateUploadUrlResponseBodyData {
	s.Folder = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetHost(v string) *GenerateUploadUrlResponseBodyData {
	s.Host = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetMessage(v string) *GenerateUploadUrlResponseBodyData {
	s.Message = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetPolicy(v string) *GenerateUploadUrlResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetSignature(v string) *GenerateUploadUrlResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) SetSuccess(v bool) *GenerateUploadUrlResponseBodyData {
	s.Success = &v
	return s
}

func (s *GenerateUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
