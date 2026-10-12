// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessType(v string) *UpdateNetworkAccessRequest
	GetNetworkAccessType() *string
}

type UpdateNetworkAccessRequest struct {
	// The expected network access type after the change. For example, to add public network access to a private network gateway, set this parameter to private&pubnet. To remove private network access from a gateway that supports both public and private network access, set this parameter to pubnet.
	//
	// Valid values:
	//
	// - privatenet: private network only.
	//
	// - private&pubnet: public network and private network.
	//
	// - pubnet: public network only.
	//
	// This parameter is required.
	//
	// example:
	//
	// private&pubnet
	NetworkAccessType *string `json:"networkAccessType,omitempty" xml:"networkAccessType,omitempty"`
}

func (s UpdateNetworkAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessRequest) GetNetworkAccessType() *string {
	return s.NetworkAccessType
}

func (s *UpdateNetworkAccessRequest) SetNetworkAccessType(v string) *UpdateNetworkAccessRequest {
	s.NetworkAccessType = &v
	return s
}

func (s *UpdateNetworkAccessRequest) Validate() error {
	return dara.Validate(s)
}
