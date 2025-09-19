// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *DeleteScalingConfigRequest
	GetQualifier() *string
}

type DeleteScalingConfigRequest struct {
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeleteScalingConfigRequest) SetQualifier(v string) *DeleteScalingConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *DeleteScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
