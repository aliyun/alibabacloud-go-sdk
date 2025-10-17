// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Template) *ListTemplateResponseBody
	GetData() []*Template
	SetErrorCode(v string) *ListTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTemplateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ListTemplateResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ListTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTemplateResponseBody
	GetSuccess() *bool
}

type ListTemplateResponseBody struct {
	Data []*Template `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ERR-00000000
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ok
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBody) GetData() []*Template {
	return s.Data
}

func (s *ListTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTemplateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTemplateResponseBody) SetData(v []*Template) *ListTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplateResponseBody) SetErrorCode(v string) *ListTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTemplateResponseBody) SetErrorMessage(v string) *ListTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTemplateResponseBody) SetHttpStatusCode(v string) *ListTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTemplateResponseBody) SetRequestId(v string) *ListTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateResponseBody) SetSuccess(v bool) *ListTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *ListTemplateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
