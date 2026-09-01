// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcCidrBlockResponseBody
	GetRequestId() *string
}

type ModifyVpcCidrBlockResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6608E72F-F276-440F-ABEF-419971CEC4D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcCidrBlockResponseBody) SetRequestId(v string) *ModifyVpcCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
