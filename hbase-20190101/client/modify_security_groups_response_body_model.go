// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityGroupsResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupsResponseBody struct {
	// example:
	//
	// F4AD2E65-482B-46B6-942E-765989B1C8A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupsResponseBody) SetRequestId(v string) *ModifySecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
