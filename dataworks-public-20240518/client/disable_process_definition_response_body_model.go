// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableProcessDefinitionResponseBody
	GetRequestId() *string
}

type DisableProcessDefinitionResponseBody struct {
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableProcessDefinitionResponseBody) SetRequestId(v string) *DisableProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableProcessDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}
