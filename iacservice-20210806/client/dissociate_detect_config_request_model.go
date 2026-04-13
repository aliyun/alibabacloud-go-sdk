// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigId(v string) *DissociateDetectConfigRequest
	GetDetectConfigId() *string
	SetTargetId(v string) *DissociateDetectConfigRequest
	GetTargetId() *string
	SetTargetType(v string) *DissociateDetectConfigRequest
	GetTargetType() *string
}

type DissociateDetectConfigRequest struct {
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stack-xxxx
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Stack
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s DissociateDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *DissociateDetectConfigRequest) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *DissociateDetectConfigRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DissociateDetectConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DissociateDetectConfigRequest) SetDetectConfigId(v string) *DissociateDetectConfigRequest {
	s.DetectConfigId = &v
	return s
}

func (s *DissociateDetectConfigRequest) SetTargetId(v string) *DissociateDetectConfigRequest {
	s.TargetId = &v
	return s
}

func (s *DissociateDetectConfigRequest) SetTargetType(v string) *DissociateDetectConfigRequest {
	s.TargetType = &v
	return s
}

func (s *DissociateDetectConfigRequest) Validate() error {
	return dara.Validate(s)
}
