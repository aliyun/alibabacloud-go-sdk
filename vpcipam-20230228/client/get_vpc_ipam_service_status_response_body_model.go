// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcIpamServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetVpcIpamServiceStatusResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *GetVpcIpamServiceStatusResponseBody
	GetRequestId() *string
}

type GetVpcIpamServiceStatusResponseBody struct {
	// Indicates whether IPAM is activated.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FEE9FFF-57EE-5832-BE88-9308352F3B68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpcIpamServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcIpamServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetVpcIpamServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcIpamServiceStatusResponseBody) SetEnabled(v bool) *GetVpcIpamServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetVpcIpamServiceStatusResponseBody) SetRequestId(v string) *GetVpcIpamServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcIpamServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
