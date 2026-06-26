// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLayerACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *PutLayerACLRequest
	GetAcl() *string
	SetPublic(v string) *PutLayerACLRequest
	GetPublic() *string
}

type PutLayerACLRequest struct {
	// The access permissions of the layer. Valid values: 1 (public) and 0 (private). The default value is 0.
	//
	// example:
	//
	// 1
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// Deprecated
	//
	// Specifies whether to make the layer public. Valid values: true and false.
	//
	// example:
	//
	// true
	Public *string `json:"public,omitempty" xml:"public,omitempty"`
}

func (s PutLayerACLRequest) String() string {
	return dara.Prettify(s)
}

func (s PutLayerACLRequest) GoString() string {
	return s.String()
}

func (s *PutLayerACLRequest) GetAcl() *string {
	return s.Acl
}

func (s *PutLayerACLRequest) GetPublic() *string {
	return s.Public
}

func (s *PutLayerACLRequest) SetAcl(v string) *PutLayerACLRequest {
	s.Acl = &v
	return s
}

func (s *PutLayerACLRequest) SetPublic(v string) *PutLayerACLRequest {
	s.Public = &v
	return s
}

func (s *PutLayerACLRequest) Validate() error {
	return dara.Validate(s)
}
