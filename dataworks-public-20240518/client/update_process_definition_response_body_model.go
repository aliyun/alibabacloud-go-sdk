// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProcessDefinitionResponseBody
	GetRequestId() *string
}

type UpdateProcessDefinitionResponseBody struct {
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProcessDefinitionResponseBody) SetRequestId(v string) *UpdateProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProcessDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
