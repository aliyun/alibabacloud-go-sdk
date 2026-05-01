// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateModelTemplateResponseBodyData) *CreateModelTemplateResponseBody
	GetData() *CreateModelTemplateResponseBodyData
	SetRequestId(v string) *CreateModelTemplateResponseBody
	GetRequestId() *string
}

type CreateModelTemplateResponseBody struct {
	Data *CreateModelTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelTemplateResponseBody) GetData() *CreateModelTemplateResponseBodyData {
	return s.Data
}

func (s *CreateModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelTemplateResponseBody) SetData(v *CreateModelTemplateResponseBodyData) *CreateModelTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelTemplateResponseBody) SetRequestId(v string) *CreateModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelTemplateResponseBodyData struct {
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
}

func (s CreateModelTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelTemplateResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *CreateModelTemplateResponseBodyData) SetModelTemplateId(v string) *CreateModelTemplateResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *CreateModelTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
