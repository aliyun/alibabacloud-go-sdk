// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContextId(v string) *DeleteContextResponseBody
	GetContextId() *string
	SetRequestId(v string) *DeleteContextResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteContextResponseBody
	GetStatus() *string
}

type DeleteContextResponseBody struct {
	// example:
	//
	// 897294a7-67a4-4f60-976c-e136edc5f97e
	ContextId *string `json:"contextId,omitempty" xml:"contextId,omitempty"`
	// example:
	//
	// E5B1D3D4-BB28-5996-8AD2-***********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteContextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextResponseBody) GetContextId() *string {
	return s.ContextId
}

func (s *DeleteContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteContextResponseBody) SetContextId(v string) *DeleteContextResponseBody {
	s.ContextId = &v
	return s
}

func (s *DeleteContextResponseBody) SetRequestId(v string) *DeleteContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextResponseBody) SetStatus(v string) *DeleteContextResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteContextResponseBody) Validate() error {
	return dara.Validate(s)
}
