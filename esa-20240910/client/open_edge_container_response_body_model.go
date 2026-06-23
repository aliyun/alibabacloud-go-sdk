// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenEdgeContainerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenEdgeContainerResponseBody
	GetRequestId() *string
}

type OpenEdgeContainerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B1C88C-9BB7-568F-9D02-DFBE6F44624D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenEdgeContainerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenEdgeContainerResponseBody) GoString() string {
	return s.String()
}

func (s *OpenEdgeContainerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenEdgeContainerResponseBody) SetRequestId(v string) *OpenEdgeContainerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenEdgeContainerResponseBody) Validate() error {
	return dara.Validate(s)
}
