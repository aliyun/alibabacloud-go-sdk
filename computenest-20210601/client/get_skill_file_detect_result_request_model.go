// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillFileDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHashKey(v string) *GetSkillFileDetectResultRequest
	GetHashKey() *string
	SetRegionId(v string) *GetSkillFileDetectResultRequest
	GetRegionId() *string
}

type GetSkillFileDetectResultRequest struct {
	// The unique identifier for the detection task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2aceb074-fa72-44d2-99d9-45b17cffe0e7
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSkillFileDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillFileDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetSkillFileDetectResultRequest) GetHashKey() *string {
	return s.HashKey
}

func (s *GetSkillFileDetectResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSkillFileDetectResultRequest) SetHashKey(v string) *GetSkillFileDetectResultRequest {
	s.HashKey = &v
	return s
}

func (s *GetSkillFileDetectResultRequest) SetRegionId(v string) *GetSkillFileDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetSkillFileDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
