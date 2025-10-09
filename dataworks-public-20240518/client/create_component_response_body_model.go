// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *CreateComponentResponseBody
	GetComponentId() *string
	SetRequestId(v string) *CreateComponentResponseBody
	GetRequestId() *string
}

type CreateComponentResponseBody struct {
	// example:
	//
	// 123123123123123
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// adssd****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComponentResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *CreateComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComponentResponseBody) SetComponentId(v string) *CreateComponentResponseBody {
	s.ComponentId = &v
	return s
}

func (s *CreateComponentResponseBody) SetRequestId(v string) *CreateComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
