// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProviderTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateModelProviderTemplateResponseBodyData) *CreateModelProviderTemplateResponseBody
	GetData() *CreateModelProviderTemplateResponseBodyData
	SetRequestId(v string) *CreateModelProviderTemplateResponseBody
	GetRequestId() *string
}

type CreateModelProviderTemplateResponseBody struct {
	Data *CreateModelProviderTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelProviderTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelProviderTemplateResponseBody) GetData() *CreateModelProviderTemplateResponseBodyData {
	return s.Data
}

func (s *CreateModelProviderTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelProviderTemplateResponseBody) SetData(v *CreateModelProviderTemplateResponseBodyData) *CreateModelProviderTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelProviderTemplateResponseBody) SetRequestId(v string) *CreateModelProviderTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelProviderTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelProviderTemplateResponseBodyData struct {
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s CreateModelProviderTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProviderTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelProviderTemplateResponseBodyData) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *CreateModelProviderTemplateResponseBodyData) SetProviderTemplateId(v string) *CreateModelProviderTemplateResponseBodyData {
	s.ProviderTemplateId = &v
	return s
}

func (s *CreateModelProviderTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
