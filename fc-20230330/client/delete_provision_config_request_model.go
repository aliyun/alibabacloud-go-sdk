// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProvisionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *DeleteProvisionConfigRequest
	GetQualifier() *string
}

type DeleteProvisionConfigRequest struct {
	// The function alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteProvisionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteProvisionConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeleteProvisionConfigRequest) SetQualifier(v string) *DeleteProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *DeleteProvisionConfigRequest) Validate() error {
	return dara.Validate(s)
}
