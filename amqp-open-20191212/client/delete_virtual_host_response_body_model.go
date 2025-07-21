// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVirtualHostResponseBody
	GetRequestId() *string
}

type DeleteVirtualHostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4311050D-BD63-48F9-822B-947A75A1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirtualHostResponseBody) SetRequestId(v string) *DeleteVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualHostResponseBody) Validate() error {
	return dara.Validate(s)
}
