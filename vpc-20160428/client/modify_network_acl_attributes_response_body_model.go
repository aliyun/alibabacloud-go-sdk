// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkAclAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNetworkAclAttributesResponseBody
	GetRequestId() *string
}

type ModifyNetworkAclAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8F76C3E4-B39F-465D-B8B3-50BAF03CA833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkAclAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkAclAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAclAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNetworkAclAttributesResponseBody) SetRequestId(v string) *ModifyNetworkAclAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNetworkAclAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
