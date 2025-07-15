// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBgpNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddBgpNetworkResponseBody
	GetRequestId() *string
}

type AddBgpNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9C7FA9D6-72E0-48A9-A9C3-2DA8569CD5EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBgpNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBgpNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *AddBgpNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBgpNetworkResponseBody) SetRequestId(v string) *AddBgpNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBgpNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
