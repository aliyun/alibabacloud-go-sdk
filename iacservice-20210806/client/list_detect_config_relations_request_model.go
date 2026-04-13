// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigId(v string) *ListDetectConfigRelationsRequest
	GetDetectConfigId() *string
	SetTargetId(v string) *ListDetectConfigRelationsRequest
	GetTargetId() *string
	SetTargetType(v string) *ListDetectConfigRelationsRequest
	GetTargetType() *string
}

type ListDetectConfigRelationsRequest struct {
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// example:
	//
	// stack-xxxxx
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// example:
	//
	// Stack
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListDetectConfigRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListDetectConfigRelationsRequest) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *ListDetectConfigRelationsRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListDetectConfigRelationsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListDetectConfigRelationsRequest) SetDetectConfigId(v string) *ListDetectConfigRelationsRequest {
	s.DetectConfigId = &v
	return s
}

func (s *ListDetectConfigRelationsRequest) SetTargetId(v string) *ListDetectConfigRelationsRequest {
	s.TargetId = &v
	return s
}

func (s *ListDetectConfigRelationsRequest) SetTargetType(v string) *ListDetectConfigRelationsRequest {
	s.TargetType = &v
	return s
}

func (s *ListDetectConfigRelationsRequest) Validate() error {
	return dara.Validate(s)
}
