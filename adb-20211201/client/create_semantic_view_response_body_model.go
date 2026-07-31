// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSemanticViewResponseBody
	GetRequestId() *string
}

type CreateSemanticViewResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSemanticViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticViewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSemanticViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSemanticViewResponseBody) SetRequestId(v string) *CreateSemanticViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSemanticViewResponseBody) Validate() error {
	return dara.Validate(s)
}
