// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneTemplateFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInterveneTemplateFileUrlResponseBody
	GetCode() *string
	SetData(v *GetInterveneTemplateFileUrlResponseBodyData) *GetInterveneTemplateFileUrlResponseBody
	GetData() *GetInterveneTemplateFileUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetInterveneTemplateFileUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInterveneTemplateFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInterveneTemplateFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInterveneTemplateFileUrlResponseBody
	GetSuccess() *bool
}

type GetInterveneTemplateFileUrlResponseBody struct {
	// example:
	//
	// 0
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneTemplateFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetData() *GetInterveneTemplateFileUrlResponseBodyData {
	return s.Data
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterveneTemplateFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetCode(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetData(v *GetInterveneTemplateFileUrlResponseBodyData) *GetInterveneTemplateFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetHttpStatusCode(v int32) *GetInterveneTemplateFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetMessage(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetRequestId(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetSuccess(v bool) *GetInterveneTemplateFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneTemplateFileUrlResponseBodyData struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) SetCode(v int32) *GetInterveneTemplateFileUrlResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) SetFileUrl(v string) *GetInterveneTemplateFileUrlResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
