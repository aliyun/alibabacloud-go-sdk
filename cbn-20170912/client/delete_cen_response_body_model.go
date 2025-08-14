// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenResponseBody
	GetRequestId() *string
}

type DeleteCenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5903EE99-D542-4E14-BC65-AAC1CB2D3D03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenResponseBody) SetRequestId(v string) *DeleteCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenResponseBody) Validate() error {
	return dara.Validate(s)
}
