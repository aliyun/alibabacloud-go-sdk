// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoNotCallFileUploadParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDoNotCallFileUploadParametersResponseBody
	GetCode() *string
	SetData(v *GetDoNotCallFileUploadParametersResponseBodyData) *GetDoNotCallFileUploadParametersResponseBody
	GetData() *GetDoNotCallFileUploadParametersResponseBodyData
	SetHttpStatusCode(v int32) *GetDoNotCallFileUploadParametersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDoNotCallFileUploadParametersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDoNotCallFileUploadParametersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDoNotCallFileUploadParametersResponseBody
	GetSuccess() *bool
}

type GetDoNotCallFileUploadParametersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDoNotCallFileUploadParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDoNotCallFileUploadParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoNotCallFileUploadParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetData() *GetDoNotCallFileUploadParametersResponseBodyData {
	return s.Data
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoNotCallFileUploadParametersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetCode(v string) *GetDoNotCallFileUploadParametersResponseBody {
	s.Code = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetData(v *GetDoNotCallFileUploadParametersResponseBodyData) *GetDoNotCallFileUploadParametersResponseBody {
	s.Data = v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetHttpStatusCode(v int32) *GetDoNotCallFileUploadParametersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetMessage(v string) *GetDoNotCallFileUploadParametersResponseBody {
	s.Message = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetRequestId(v string) *GetDoNotCallFileUploadParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) SetSuccess(v bool) *GetDoNotCallFileUploadParametersResponseBody {
	s.Success = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoNotCallFileUploadParametersResponseBodyData struct {
	// example:
	//
	// ****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 1647313420
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// ccc-test/blacklist.xlsx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMi0wMy0xNVQwMzowMzo0MC4zMTJaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjBdLFsic3RhcnRzLXdpdGgiLCIka2V5IiwidGVtcC9ibGFja2xpc3QvaW1wb3J0LzE1NzcyNDcxMTU0OTA0MDEvY2NjVjIta216LzIwMjIE1MTAwMzQwLyJd****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// zi31STIMtIfa/UN2l+6lww****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetDoNotCallFileUploadParametersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoNotCallFileUploadParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetAccessKeyId(v string) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetExpireTime(v int32) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetFilePath(v string) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetHost(v string) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetPolicy(v string) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) SetSignature(v string) *GetDoNotCallFileUploadParametersResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
