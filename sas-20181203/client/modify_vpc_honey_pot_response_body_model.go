// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcHoneyPotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcHoneyPotResponseBody
	GetRequestId() *string
}

type ModifyVpcHoneyPotResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8158FE9E-19BE-42D6-9F7A-9BE34A2DE1D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcHoneyPotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcHoneyPotResponseBody) SetRequestId(v string) *ModifyVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcHoneyPotResponseBody) Validate() error {
	return dara.Validate(s)
}
