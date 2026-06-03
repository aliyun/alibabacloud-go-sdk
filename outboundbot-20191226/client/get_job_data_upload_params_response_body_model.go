// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDataUploadParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetJobDataUploadParamsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetJobDataUploadParamsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetJobDataUploadParamsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobDataUploadParamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobDataUploadParamsResponseBody
	GetSuccess() *bool
	SetUploadParams(v *GetJobDataUploadParamsResponseBodyUploadParams) *GetJobDataUploadParamsResponseBody
	GetUploadParams() *GetJobDataUploadParamsResponseBodyUploadParams
}

type GetJobDataUploadParamsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// True
	Success      *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	UploadParams *GetJobDataUploadParamsResponseBodyUploadParams `json:"UploadParams,omitempty" xml:"UploadParams,omitempty" type:"Struct"`
}

func (s GetJobDataUploadParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobDataUploadParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDataUploadParamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJobDataUploadParamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetJobDataUploadParamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobDataUploadParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobDataUploadParamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobDataUploadParamsResponseBody) GetUploadParams() *GetJobDataUploadParamsResponseBodyUploadParams {
	return s.UploadParams
}

func (s *GetJobDataUploadParamsResponseBody) SetCode(v string) *GetJobDataUploadParamsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) SetHttpStatusCode(v int32) *GetJobDataUploadParamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) SetMessage(v string) *GetJobDataUploadParamsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) SetRequestId(v string) *GetJobDataUploadParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) SetSuccess(v bool) *GetJobDataUploadParamsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) SetUploadParams(v *GetJobDataUploadParamsResponseBodyUploadParams) *GetJobDataUploadParamsResponseBody {
	s.UploadParams = v
	return s
}

func (s *GetJobDataUploadParamsResponseBody) Validate() error {
	if s.UploadParams != nil {
		if err := s.UploadParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobDataUploadParamsResponseBodyUploadParams struct {
	// example:
	//
	// LTAIvKWEr4DoFSqz
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 1741855527
	Expire *int32 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// UPLOADED/SCRIPT/136a055e-3d07-46c6-853a-731b3a2973de/18dc6d79-338f-413c-a5a8-dcce5f05b41a_9bd7916d-a340-4925-a911-92390cbe0f33.json
	Folder *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	// example:
	//
	// https://cloudagentbot-online.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0xM1QwODo0NToyNy4zMzFaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNTI0Mjg4MDBdLFsic3RhcnRzLXdpdGgiLCIka2V5IiwiVVBMT0FERUQvU0NSSVBULzEzNmEwNTVlLTNkMDctNDZjNi04NTNhLTczMWIzYTI5NzNkZS8iXV19
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// MD4CHQCiECtjdsP+fstJDcqsPt+GbWxSWPzGQxeQiblzAh0AuidaKc5JY5AkHFIE+qiQI3izJQQbpUoG0paPTw==
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetJobDataUploadParamsResponseBodyUploadParams) String() string {
	return dara.Prettify(s)
}

func (s GetJobDataUploadParamsResponseBodyUploadParams) GoString() string {
	return s.String()
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetAccessId() *string {
	return s.AccessId
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetExpire() *int32 {
	return s.Expire
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetFolder() *string {
	return s.Folder
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetHost() *string {
	return s.Host
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetPolicy() *string {
	return s.Policy
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) GetSignature() *string {
	return s.Signature
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetAccessId(v string) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.AccessId = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetExpire(v int32) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.Expire = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetFolder(v string) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.Folder = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetHost(v string) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.Host = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetPolicy(v string) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.Policy = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) SetSignature(v string) *GetJobDataUploadParamsResponseBodyUploadParams {
	s.Signature = &v
	return s
}

func (s *GetJobDataUploadParamsResponseBodyUploadParams) Validate() error {
	return dara.Validate(s)
}
