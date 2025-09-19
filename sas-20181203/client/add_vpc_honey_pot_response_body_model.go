// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVpcHoneyPotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVpcHoneyPotResponseBody
	GetRequestId() *string
}

type AddVpcHoneyPotResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVpcHoneyPotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVpcHoneyPotResponseBody) SetRequestId(v string) *AddVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVpcHoneyPotResponseBody) Validate() error {
	return dara.Validate(s)
}
