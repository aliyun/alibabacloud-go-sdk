// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDefenseResourceGroupResponseBody
	GetRequestId() *string
}

type DeleteDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25BE1169-9AE4-5D7D-8293-C33242ABB549
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDefenseResourceGroupResponseBody) SetRequestId(v string) *DeleteDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDefenseResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
