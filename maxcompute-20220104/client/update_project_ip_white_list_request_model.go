// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectIpWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateProjectIpWhiteListRequest
	GetBody() *string
}

type UpdateProjectIpWhiteListRequest struct {
	// The request body parameters.
	//
	// example:
	//
	// json {"ipWhiteList":{"ipList": "", // The IP address whitelists are of the STRING data type. Separate multiple IP address whitelists with commas (,). "vpcIpList": "", //} }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectIpWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateProjectIpWhiteListRequest) SetBody(v string) *UpdateProjectIpWhiteListRequest {
	s.Body = &v
	return s
}

func (s *UpdateProjectIpWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
