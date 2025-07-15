// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicIpAddressPoolServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *GetPublicIpAddressPoolServiceStatusResponseBody
	GetEnabled() *bool
	SetRequestId(v string) *GetPublicIpAddressPoolServiceStatusResponseBody
	GetRequestId() *string
}

type GetPublicIpAddressPoolServiceStatusResponseBody struct {
	// Indicates whether the IP address pool feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- You can call OpenPublicIpAddressPoolService to enable the IP address pool feature.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-06F82A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicIpAddressPoolServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublicIpAddressPoolServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicIpAddressPoolServiceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetPublicIpAddressPoolServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublicIpAddressPoolServiceStatusResponseBody) SetEnabled(v bool) *GetPublicIpAddressPoolServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusResponseBody) SetRequestId(v string) *GetPublicIpAddressPoolServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
