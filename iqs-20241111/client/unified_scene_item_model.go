// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSceneItem interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *UnifiedSceneItem
	GetDetail() *string
	SetType(v string) *UnifiedSceneItem
	GetType() *string
}

type UnifiedSceneItem struct {
	// The result in JSON structure.
	//
	// example:
	//
	// {\\"title\\": \\"伦敦时间\\", \\"targetTimeZone\\": \\"Europe/London\\", \\"targetTimeMillisecond\\": \\"1745820817178\\", \\"targetTime\\": \\"2025-04-28 14:13:37\\", \\"beijingTimeZone\\": \\"PRC\\", \\"beijingTimeMillisecond\\": \\"1745846017178\\"}
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// The type.
	//
	// example:
	//
	// time
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UnifiedSceneItem) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSceneItem) GoString() string {
	return s.String()
}

func (s *UnifiedSceneItem) GetDetail() *string {
	return s.Detail
}

func (s *UnifiedSceneItem) GetType() *string {
	return s.Type
}

func (s *UnifiedSceneItem) SetDetail(v string) *UnifiedSceneItem {
	s.Detail = &v
	return s
}

func (s *UnifiedSceneItem) SetType(v string) *UnifiedSceneItem {
	s.Type = &v
	return s
}

func (s *UnifiedSceneItem) Validate() error {
	return dara.Validate(s)
}
