// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigId(v string) *AssociateDetectConfigRequest
	GetDetectConfigId() *string
	SetTargetId(v string) *AssociateDetectConfigRequest
	GetTargetId() *string
	SetTargetType(v string) *AssociateDetectConfigRequest
	GetTargetType() *string
}

type AssociateDetectConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stack-xxxxx
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Stack
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s AssociateDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *AssociateDetectConfigRequest) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *AssociateDetectConfigRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *AssociateDetectConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *AssociateDetectConfigRequest) SetDetectConfigId(v string) *AssociateDetectConfigRequest {
	s.DetectConfigId = &v
	return s
}

func (s *AssociateDetectConfigRequest) SetTargetId(v string) *AssociateDetectConfigRequest {
	s.TargetId = &v
	return s
}

func (s *AssociateDetectConfigRequest) SetTargetType(v string) *AssociateDetectConfigRequest {
	s.TargetType = &v
	return s
}

func (s *AssociateDetectConfigRequest) Validate() error {
	return dara.Validate(s)
}
