// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebFilterTemplateId(v string) *DeleteWebFilterTemplateRequest
	GetWebFilterTemplateId() *string
}

type DeleteWebFilterTemplateRequest struct {
	// The ID of the web filtering template to delete. This parameter is required. If you leave this parameter empty, a parameter fault is returned. The template must exist within the current Alibaba Cloud account. A template cannot be deleted if it is referenced by an access control policy of the Internet border, NAT firewalls, or virtual private cloud (VPC) firewalls.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s DeleteWebFilterTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterTemplateRequest) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *DeleteWebFilterTemplateRequest) SetWebFilterTemplateId(v string) *DeleteWebFilterTemplateRequest {
	s.WebFilterTemplateId = &v
	return s
}

func (s *DeleteWebFilterTemplateRequest) Validate() error {
	return dara.Validate(s)
}
