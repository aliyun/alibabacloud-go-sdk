// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSemanticViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSemanticViewResponseBody
	GetRequestId() *string
}

type DeleteSemanticViewResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3A8F6106-6AFD-5A34-9C80-8DE2C42D06E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSemanticViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSemanticViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSemanticViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSemanticViewResponseBody) SetRequestId(v string) *DeleteSemanticViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSemanticViewResponseBody) Validate() error {
	return dara.Validate(s)
}
