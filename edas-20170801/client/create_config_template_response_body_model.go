// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateConfigTemplateResponseBody
	GetCode() *int32
	SetData(v *CreateConfigTemplateResponseBodyData) *CreateConfigTemplateResponseBody
	GetData() *CreateConfigTemplateResponseBodyData
	SetMessage(v string) *CreateConfigTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConfigTemplateResponseBody
	GetRequestId() *string
}

type CreateConfigTemplateResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateConfigTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateConfigTemplateResponseBody) GetData() *CreateConfigTemplateResponseBodyData {
	return s.Data
}

func (s *CreateConfigTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConfigTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigTemplateResponseBody) SetCode(v int32) *CreateConfigTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetData(v *CreateConfigTemplateResponseBodyData) *CreateConfigTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetMessage(v string) *CreateConfigTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigTemplateResponseBody) SetRequestId(v string) *CreateConfigTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConfigTemplateResponseBodyData struct {
	// The ID of the template.
	//
	// example:
	//
	// 125122
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateConfigTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CreateConfigTemplateResponseBodyData) SetId(v int64) *CreateConfigTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateConfigTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
