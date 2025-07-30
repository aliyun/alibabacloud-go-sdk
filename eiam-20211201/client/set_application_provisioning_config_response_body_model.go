// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationProvisioningConfigResponseBody
	GetRequestId() *string
}

type SetApplicationProvisioningConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationProvisioningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationProvisioningConfigResponseBody) SetRequestId(v string) *SetApplicationProvisioningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationProvisioningConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
