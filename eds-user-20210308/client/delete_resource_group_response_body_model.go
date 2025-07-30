// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceGroupResponseBody
	GetRequestId() *string
}

type DeleteResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2463A343-BD32-5803-959E-9A8472A1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
