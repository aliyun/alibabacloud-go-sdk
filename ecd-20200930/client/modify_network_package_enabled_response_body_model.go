// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNetworkPackageEnabledResponseBody
	GetRequestId() *string
}

type ModifyNetworkPackageEnabledResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNetworkPackageEnabledResponseBody) SetRequestId(v string) *ModifyNetworkPackageEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
