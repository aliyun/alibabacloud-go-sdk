// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiRequestContract interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *HttpApiRequestContractBody) *HttpApiRequestContract
	GetBody() *HttpApiRequestContractBody
	SetHeaderParameters(v []*HttpApiParameter) *HttpApiRequestContract
	GetHeaderParameters() []*HttpApiParameter
	SetPathParameters(v []*HttpApiParameter) *HttpApiRequestContract
	GetPathParameters() []*HttpApiParameter
	SetQueryParameters(v []*HttpApiParameter) *HttpApiRequestContract
	GetQueryParameters() []*HttpApiParameter
}

type HttpApiRequestContract struct {
	// The body parameters.
	Body *HttpApiRequestContractBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The request header parameters.
	HeaderParameters []*HttpApiParameter `json:"headerParameters,omitempty" xml:"headerParameters,omitempty" type:"Repeated"`
	// The path parameters.
	PathParameters []*HttpApiParameter `json:"pathParameters,omitempty" xml:"pathParameters,omitempty" type:"Repeated"`
	// The query parameters.
	QueryParameters []*HttpApiParameter `json:"queryParameters,omitempty" xml:"queryParameters,omitempty" type:"Repeated"`
}

func (s HttpApiRequestContract) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRequestContract) GoString() string {
	return s.String()
}

func (s *HttpApiRequestContract) GetBody() *HttpApiRequestContractBody {
	return s.Body
}

func (s *HttpApiRequestContract) GetHeaderParameters() []*HttpApiParameter {
	return s.HeaderParameters
}

func (s *HttpApiRequestContract) GetPathParameters() []*HttpApiParameter {
	return s.PathParameters
}

func (s *HttpApiRequestContract) GetQueryParameters() []*HttpApiParameter {
	return s.QueryParameters
}

func (s *HttpApiRequestContract) SetBody(v *HttpApiRequestContractBody) *HttpApiRequestContract {
	s.Body = v
	return s
}

func (s *HttpApiRequestContract) SetHeaderParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.HeaderParameters = v
	return s
}

func (s *HttpApiRequestContract) SetPathParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.PathParameters = v
	return s
}

func (s *HttpApiRequestContract) SetQueryParameters(v []*HttpApiParameter) *HttpApiRequestContract {
	s.QueryParameters = v
	return s
}

func (s *HttpApiRequestContract) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderParameters != nil {
		for _, item := range s.HeaderParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PathParameters != nil {
		for _, item := range s.PathParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryParameters != nil {
		for _, item := range s.QueryParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiRequestContractBody struct {
	// The content type of the request body.
	//
	// example:
	//
	// application/json
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// This is a description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The sample value.
	//
	// example:
	//
	// {"result":"ok"}
	Example *string `json:"example,omitempty" xml:"example,omitempty"`
	// The JSON definition description of the request body.
	//
	// example:
	//
	// {
	//
	//       "type": "object",
	//
	//       "required": [
	//
	//           "result"
	//
	//       ],
	//
	//       "properties": {
	//
	//           "result": {
	//
	//               "type": "string",
	//
	//               "description": "Operation result. \\"ok\\" indicates success."
	//
	//           }
	//
	//       }
	//
	//   }
	JsonSchema *string `json:"jsonSchema,omitempty" xml:"jsonSchema,omitempty"`
}

func (s HttpApiRequestContractBody) String() string {
	return dara.Prettify(s)
}

func (s HttpApiRequestContractBody) GoString() string {
	return s.String()
}

func (s *HttpApiRequestContractBody) GetContentType() *string {
	return s.ContentType
}

func (s *HttpApiRequestContractBody) GetDescription() *string {
	return s.Description
}

func (s *HttpApiRequestContractBody) GetExample() *string {
	return s.Example
}

func (s *HttpApiRequestContractBody) GetJsonSchema() *string {
	return s.JsonSchema
}

func (s *HttpApiRequestContractBody) SetContentType(v string) *HttpApiRequestContractBody {
	s.ContentType = &v
	return s
}

func (s *HttpApiRequestContractBody) SetDescription(v string) *HttpApiRequestContractBody {
	s.Description = &v
	return s
}

func (s *HttpApiRequestContractBody) SetExample(v string) *HttpApiRequestContractBody {
	s.Example = &v
	return s
}

func (s *HttpApiRequestContractBody) SetJsonSchema(v string) *HttpApiRequestContractBody {
	s.JsonSchema = &v
	return s
}

func (s *HttpApiRequestContractBody) Validate() error {
	return dara.Validate(s)
}
