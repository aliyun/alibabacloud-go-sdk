// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutScalingConfigInput) *PutScalingConfigRequest
	GetBody() *PutScalingConfigInput
	SetQualifier(v string) *PutScalingConfigRequest
	GetQualifier() *string
}

type PutScalingConfigRequest struct {
	Body *PutScalingConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *PutScalingConfigRequest) GetBody() *PutScalingConfigInput {
	return s.Body
}

func (s *PutScalingConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *PutScalingConfigRequest) SetBody(v *PutScalingConfigInput) *PutScalingConfigRequest {
	s.Body = v
	return s
}

func (s *PutScalingConfigRequest) SetQualifier(v string) *PutScalingConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *PutScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
