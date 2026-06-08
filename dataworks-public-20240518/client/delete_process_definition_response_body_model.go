// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProcessDefinitionResponseBody
	GetRequestId() *string
}

type DeleteProcessDefinitionResponseBody struct {
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProcessDefinitionResponseBody) SetRequestId(v string) *DeleteProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProcessDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
