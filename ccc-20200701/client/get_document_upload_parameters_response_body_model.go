// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUploadParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentUploadParametersResponseBody
	GetCode() *string
	SetData(v *GetDocumentUploadParametersResponseBodyData) *GetDocumentUploadParametersResponseBody
	GetData() *GetDocumentUploadParametersResponseBodyData
	SetHttpStatusCode(v int32) *GetDocumentUploadParametersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocumentUploadParametersResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetDocumentUploadParametersResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetDocumentUploadParametersResponseBody
	GetRequestId() *string
}

type GetDocumentUploadParametersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocumentUploadParametersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9FBA26B0-462B-4D77-B78F-AF35560DBC71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDocumentUploadParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUploadParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentUploadParametersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentUploadParametersResponseBody) GetData() *GetDocumentUploadParametersResponseBodyData {
	return s.Data
}

func (s *GetDocumentUploadParametersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocumentUploadParametersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentUploadParametersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetDocumentUploadParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentUploadParametersResponseBody) SetCode(v string) *GetDocumentUploadParametersResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) SetData(v *GetDocumentUploadParametersResponseBodyData) *GetDocumentUploadParametersResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) SetHttpStatusCode(v int32) *GetDocumentUploadParametersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) SetMessage(v string) *GetDocumentUploadParametersResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) SetParams(v []*string) *GetDocumentUploadParametersResponseBody {
	s.Params = v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) SetRequestId(v string) *GetDocumentUploadParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocumentUploadParametersResponseBodyData struct {
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
	// oss host
	//
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// zi31STIMtIfa/UN2l+6lww****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetDocumentUploadParametersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUploadParametersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentUploadParametersResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetDocumentUploadParametersResponseBodyData) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GetDocumentUploadParametersResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetDocumentUploadParametersResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetDocumentUploadParametersResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetDocumentUploadParametersResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetDocumentUploadParametersResponseBodyData) SetAccessKeyId(v string) *GetDocumentUploadParametersResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) SetExpireTime(v int32) *GetDocumentUploadParametersResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) SetFilePath(v string) *GetDocumentUploadParametersResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) SetHost(v string) *GetDocumentUploadParametersResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) SetPolicy(v string) *GetDocumentUploadParametersResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) SetSignature(v string) *GetDocumentUploadParametersResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetDocumentUploadParametersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
