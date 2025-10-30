// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListSchemasResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSchemasResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSchemasResponseBody
	GetRequestId() *string
	SetSchemas(v *ListSchemasResponseBodySchemas) *ListSchemasResponseBody
	GetSchemas() *ListSchemasResponseBodySchemas
	SetStatus(v string) *ListSchemasResponseBody
	GetStatus() *string
}

type ListSchemasResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried schemas.
	Schemas *ListSchemasResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSchemasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemasResponseBody) GetSchemas() *ListSchemasResponseBodySchemas {
	return s.Schemas
}

func (s *ListSchemasResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListSchemasResponseBody) SetMessage(v string) *ListSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *ListSchemasResponseBody) SetNextToken(v string) *ListSchemasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSchemasResponseBody) SetRequestId(v string) *ListSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemasResponseBody) SetSchemas(v *ListSchemasResponseBodySchemas) *ListSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *ListSchemasResponseBody) SetStatus(v string) *ListSchemasResponseBody {
	s.Status = &v
	return s
}

func (s *ListSchemasResponseBody) Validate() error {
	if s.Schemas != nil {
		if err := s.Schemas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemasResponseBodySchemas struct {
	Schemas []*string `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s ListSchemasResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBodySchemas) GetSchemas() []*string {
	return s.Schemas
}

func (s *ListSchemasResponseBodySchemas) SetSchemas(v []*string) *ListSchemasResponseBodySchemas {
	s.Schemas = v
	return s
}

func (s *ListSchemasResponseBodySchemas) Validate() error {
	return dara.Validate(s)
}
