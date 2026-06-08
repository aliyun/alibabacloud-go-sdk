// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateProcessDefinitionResponseBody
	GetId() *string
	SetRequestId(v string) *CreateProcessDefinitionResponseBody
	GetRequestId() *string
}

type CreateProcessDefinitionResponseBody struct {
	// example:
	//
	// 1010543619
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProcessDefinitionResponseBody) SetId(v string) *CreateProcessDefinitionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateProcessDefinitionResponseBody) SetRequestId(v string) *CreateProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProcessDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
