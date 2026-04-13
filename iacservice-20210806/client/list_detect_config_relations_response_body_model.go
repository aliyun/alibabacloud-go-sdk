// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigRelations(v []*ListDetectConfigRelationsResponseBodyDetectConfigRelations) *ListDetectConfigRelationsResponseBody
	GetDetectConfigRelations() []*ListDetectConfigRelationsResponseBodyDetectConfigRelations
	SetRequestId(v string) *ListDetectConfigRelationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDetectConfigRelationsResponseBody
	GetTotalCount() *int32
}

type ListDetectConfigRelationsResponseBody struct {
	DetectConfigRelations []*ListDetectConfigRelationsResponseBodyDetectConfigRelations `json:"detectConfigRelations,omitempty" xml:"detectConfigRelations,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-XXXXXXXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDetectConfigRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectConfigRelationsResponseBody) GetDetectConfigRelations() []*ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	return s.DetectConfigRelations
}

func (s *ListDetectConfigRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDetectConfigRelationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDetectConfigRelationsResponseBody) SetDetectConfigRelations(v []*ListDetectConfigRelationsResponseBodyDetectConfigRelations) *ListDetectConfigRelationsResponseBody {
	s.DetectConfigRelations = v
	return s
}

func (s *ListDetectConfigRelationsResponseBody) SetRequestId(v string) *ListDetectConfigRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBody) SetTotalCount(v int32) *ListDetectConfigRelationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBody) Validate() error {
	if s.DetectConfigRelations != nil {
		for _, item := range s.DetectConfigRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDetectConfigRelationsResponseBodyDetectConfigRelations struct {
	// example:
	//
	// 2026-04-08T08:53:07.000+00:00
	AttachDate *string `json:"attachDate,omitempty" xml:"attachDate,omitempty"`
	// example:
	//
	// dc-xxxx
	DetectConfigId *string `json:"detectConfigId,omitempty" xml:"detectConfigId,omitempty"`
	// example:
	//
	// true
	Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// stack-xxxx
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// example:
	//
	// stack-name-xxx
	TargetName *string `json:"targetName,omitempty" xml:"targetName,omitempty"`
	// example:
	//
	// Stack
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListDetectConfigRelationsResponseBodyDetectConfigRelations) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigRelationsResponseBodyDetectConfigRelations) GoString() string {
	return s.String()
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetDetectConfigId() *string {
	return s.DetectConfigId
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetEnabled() *string {
	return s.Enabled
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetTargetId() *string {
	return s.TargetId
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetTargetName() *string {
	return s.TargetName
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) GetTargetType() *string {
	return s.TargetType
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetAttachDate(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.AttachDate = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetDetectConfigId(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.DetectConfigId = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetEnabled(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.Enabled = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetTargetId(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.TargetId = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetTargetName(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.TargetName = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) SetTargetType(v string) *ListDetectConfigRelationsResponseBodyDetectConfigRelations {
	s.TargetType = &v
	return s
}

func (s *ListDetectConfigRelationsResponseBodyDetectConfigRelations) Validate() error {
	return dara.Validate(s)
}
