// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseFileUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCaseFileUploadUrlResponseBody
	GetCode() *string
	SetData(v *GetCaseFileUploadUrlResponseBodyData) *GetCaseFileUploadUrlResponseBody
	GetData() *GetCaseFileUploadUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetCaseFileUploadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCaseFileUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCaseFileUploadUrlResponseBody
	GetRequestId() *string
}

type GetCaseFileUploadUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCaseFileUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaseFileUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCaseFileUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaseFileUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCaseFileUploadUrlResponseBody) GetData() *GetCaseFileUploadUrlResponseBodyData {
	return s.Data
}

func (s *GetCaseFileUploadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCaseFileUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCaseFileUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCaseFileUploadUrlResponseBody) SetCode(v string) *GetCaseFileUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBody) SetData(v *GetCaseFileUploadUrlResponseBodyData) *GetCaseFileUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetCaseFileUploadUrlResponseBody) SetHttpStatusCode(v int32) *GetCaseFileUploadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBody) SetMessage(v string) *GetCaseFileUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBody) SetRequestId(v string) *GetCaseFileUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCaseFileUploadUrlResponseBodyData struct {
	// example:
	//
	// ccc-test/namelist.csv
	CaseFileKey *string `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	// example:
	//
	// https://ccc-v2-online.oss-cn-shanghai.aliyuncs.com/ccc-test/namelist.csv?Expires=1642067227&OSSAccessKeyId=****&Signature=****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetCaseFileUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCaseFileUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCaseFileUploadUrlResponseBodyData) GetCaseFileKey() *string {
	return s.CaseFileKey
}

func (s *GetCaseFileUploadUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetCaseFileUploadUrlResponseBodyData) SetCaseFileKey(v string) *GetCaseFileUploadUrlResponseBodyData {
	s.CaseFileKey = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBodyData) SetUrl(v string) *GetCaseFileUploadUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetCaseFileUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
