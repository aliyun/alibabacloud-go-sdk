// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiResponseContract interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *HttpApiResponseContract
	GetContentType() *string
	SetItems(v []*HttpApiResponseContractItems) *HttpApiResponseContract
	GetItems() []*HttpApiResponseContractItems
}

type HttpApiResponseContract struct {
	// The content type.
	//
	// This parameter is required.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The response definition.
	Items []*HttpApiResponseContractItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s HttpApiResponseContract) String() string {
	return dara.Prettify(s)
}

func (s HttpApiResponseContract) GoString() string {
	return s.String()
}

func (s *HttpApiResponseContract) GetContentType() *string {
	return s.ContentType
}

func (s *HttpApiResponseContract) GetItems() []*HttpApiResponseContractItems {
	return s.Items
}

func (s *HttpApiResponseContract) SetContentType(v string) *HttpApiResponseContract {
	s.ContentType = &v
	return s
}

func (s *HttpApiResponseContract) SetItems(v []*HttpApiResponseContractItems) *HttpApiResponseContract {
	s.Items = v
	return s
}

func (s *HttpApiResponseContract) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiResponseContractItems struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The response description.
	//
	// example:
	//
	// This is a description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The sample value.
	//
	// example:
	//
	// {"result": "ok"}
	Example *string `json:"example,omitempty" xml:"example,omitempty"`
	// The JSON definition description of the response body.
	//
	// example:
	//
	// {
	//
	//     "type": "object",
	//
	//     "required": [
	//
	//         "result"
	//
	//     ],
	//
	//     "properties": {
	//
	//         "result": {
	//
	//             "type": "string",
	//
	//             "description": "This is a description."
	//
	//         }
	//
	//     }
	//
	// }
	JsonSchema *string `json:"jsonSchema,omitempty" xml:"jsonSchema,omitempty"`
}

func (s HttpApiResponseContractItems) String() string {
	return dara.Prettify(s)
}

func (s HttpApiResponseContractItems) GoString() string {
	return s.String()
}

func (s *HttpApiResponseContractItems) GetCode() *int32 {
	return s.Code
}

func (s *HttpApiResponseContractItems) GetDescription() *string {
	return s.Description
}

func (s *HttpApiResponseContractItems) GetExample() *string {
	return s.Example
}

func (s *HttpApiResponseContractItems) GetJsonSchema() *string {
	return s.JsonSchema
}

func (s *HttpApiResponseContractItems) SetCode(v int32) *HttpApiResponseContractItems {
	s.Code = &v
	return s
}

func (s *HttpApiResponseContractItems) SetDescription(v string) *HttpApiResponseContractItems {
	s.Description = &v
	return s
}

func (s *HttpApiResponseContractItems) SetExample(v string) *HttpApiResponseContractItems {
	s.Example = &v
	return s
}

func (s *HttpApiResponseContractItems) SetJsonSchema(v string) *HttpApiResponseContractItems {
	s.JsonSchema = &v
	return s
}

func (s *HttpApiResponseContractItems) Validate() error {
	return dara.Validate(s)
}
