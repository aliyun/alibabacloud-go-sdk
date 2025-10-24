// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyTemplateModel(v *GetInstancePropertiesResponseBodyPropertyTemplateModel) *GetInstancePropertiesResponseBody
	GetPropertyTemplateModel() *GetInstancePropertiesResponseBodyPropertyTemplateModel
	SetRequestId(v string) *GetInstancePropertiesResponseBody
	GetRequestId() *string
}

type GetInstancePropertiesResponseBody struct {
	PropertyTemplateModel *GetInstancePropertiesResponseBodyPropertyTemplateModel `json:"PropertyTemplateModel,omitempty" xml:"PropertyTemplateModel,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstancePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancePropertiesResponseBody) GetPropertyTemplateModel() *GetInstancePropertiesResponseBodyPropertyTemplateModel {
	return s.PropertyTemplateModel
}

func (s *GetInstancePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstancePropertiesResponseBody) SetPropertyTemplateModel(v *GetInstancePropertiesResponseBodyPropertyTemplateModel) *GetInstancePropertiesResponseBody {
	s.PropertyTemplateModel = v
	return s
}

func (s *GetInstancePropertiesResponseBody) SetRequestId(v string) *GetInstancePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstancePropertiesResponseBody) Validate() error {
	if s.PropertyTemplateModel != nil {
		if err := s.PropertyTemplateModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancePropertiesResponseBodyPropertyTemplateModel struct {
	// example:
	//
	// { "propties":{"key1":"value1"}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetInstancePropertiesResponseBodyPropertyTemplateModel) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePropertiesResponseBodyPropertyTemplateModel) GoString() string {
	return s.String()
}

func (s *GetInstancePropertiesResponseBodyPropertyTemplateModel) GetContent() *string {
	return s.Content
}

func (s *GetInstancePropertiesResponseBodyPropertyTemplateModel) SetContent(v string) *GetInstancePropertiesResponseBodyPropertyTemplateModel {
	s.Content = &v
	return s
}

func (s *GetInstancePropertiesResponseBodyPropertyTemplateModel) Validate() error {
	return dara.Validate(s)
}
