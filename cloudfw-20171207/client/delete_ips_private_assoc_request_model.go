// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsPrivateAssocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteIpsPrivateAssocRequest
	GetLang() *string
	SetResourceId(v string) *DeleteIpsPrivateAssocRequest
	GetResourceId() *string
}

type DeleteIpsPrivateAssocRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// ngw-c5vhmjdfp5t****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DeleteIpsPrivateAssocRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsPrivateAssocRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpsPrivateAssocRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteIpsPrivateAssocRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DeleteIpsPrivateAssocRequest) SetLang(v string) *DeleteIpsPrivateAssocRequest {
	s.Lang = &v
	return s
}

func (s *DeleteIpsPrivateAssocRequest) SetResourceId(v string) *DeleteIpsPrivateAssocRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteIpsPrivateAssocRequest) Validate() error {
	return dara.Validate(s)
}
