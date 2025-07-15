// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileUploadParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAudioFileUploadParametersResponseBody
	GetCode() *string
	SetData(v *GetAudioFileUploadParametersResponseBodyData) *GetAudioFileUploadParametersResponseBody
	GetData() *GetAudioFileUploadParametersResponseBodyData
	SetHttpStatusCode(v int32) *GetAudioFileUploadParametersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAudioFileUploadParametersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAudioFileUploadParametersResponseBody
	GetRequestId() *string
}

type GetAudioFileUploadParametersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAudioFileUploadParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C81FD1A5-4B99-470A-A527-D80150228784
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAudioFileUploadParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileUploadParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioFileUploadParametersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAudioFileUploadParametersResponseBody) GetData() *GetAudioFileUploadParametersResponseBodyData {
	return s.Data
}

func (s *GetAudioFileUploadParametersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAudioFileUploadParametersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAudioFileUploadParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAudioFileUploadParametersResponseBody) SetCode(v string) *GetAudioFileUploadParametersResponseBody {
	s.Code = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBody) SetData(v *GetAudioFileUploadParametersResponseBodyData) *GetAudioFileUploadParametersResponseBody {
	s.Data = v
	return s
}

func (s *GetAudioFileUploadParametersResponseBody) SetHttpStatusCode(v int32) *GetAudioFileUploadParametersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBody) SetMessage(v string) *GetAudioFileUploadParametersResponseBody {
	s.Message = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBody) SetRequestId(v string) *GetAudioFileUploadParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAudioFileUploadParametersResponseBodyData struct {
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
	// ccc-test/test-file.wav
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMi0wNi0yOVQwMDoyOTowMC42NTRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjBdLFsic3RhcnRzLXdpdGgiLCIka2V5IiwiYXVkaW8vMTU3NzI0NzExNTQ5MDQwMS9seS1vbmxpbmUvMjAyMjA2MjkwNzI5MDAvIl1d****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// HIyClras8IcVlbTV7RIJWJbU****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetAudioFileUploadParametersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileUploadParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetAudioFileUploadParametersResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetAccessKeyId(v string) *GetAudioFileUploadParametersResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetExpireTime(v int32) *GetAudioFileUploadParametersResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetFilePath(v string) *GetAudioFileUploadParametersResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetHost(v string) *GetAudioFileUploadParametersResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetPolicy(v string) *GetAudioFileUploadParametersResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) SetSignature(v string) *GetAudioFileUploadParametersResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetAudioFileUploadParametersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
