// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddNetworkInterfaceToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddNetworkInterfaceToInstanceResponseBody
	GetRequestId() *string
}

type AddNetworkInterfaceToInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddNetworkInterfaceToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddNetworkInterfaceToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddNetworkInterfaceToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddNetworkInterfaceToInstanceResponseBody) SetRequestId(v string) *AddNetworkInterfaceToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNetworkInterfaceToInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
