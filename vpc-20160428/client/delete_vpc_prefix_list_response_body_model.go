// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcPrefixListResponseBody
	GetRequestId() *string
}

type DeleteVpcPrefixListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 64B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcPrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcPrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcPrefixListResponseBody) SetRequestId(v string) *DeleteVpcPrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcPrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
