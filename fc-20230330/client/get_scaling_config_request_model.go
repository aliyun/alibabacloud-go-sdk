// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetScalingConfigRequest
	GetQualifier() *string
}

type GetScalingConfigRequest struct {
	// The alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *GetScalingConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetScalingConfigRequest) SetQualifier(v string) *GetScalingConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *GetScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
