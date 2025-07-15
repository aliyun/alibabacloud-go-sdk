// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterVideoResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCasterVideoResourceResponseBody
	GetRequestId() *string
}

type DeleteCasterVideoResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122C*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterVideoResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterVideoResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterVideoResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterVideoResourceResponseBody) SetRequestId(v string) *DeleteCasterVideoResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterVideoResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
