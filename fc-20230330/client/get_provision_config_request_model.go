// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetProvisionConfigRequest
	GetQualifier() *string
}

type GetProvisionConfigRequest struct {
	// The function alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetProvisionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetProvisionConfigRequest) SetQualifier(v string) *GetProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *GetProvisionConfigRequest) Validate() error {
	return dara.Validate(s)
}
