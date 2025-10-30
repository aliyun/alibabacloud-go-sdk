// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListEvaluationMetadataRequest
	GetLanguage() *string
	SetLensCode(v string) *ListEvaluationMetadataRequest
	GetLensCode() *string
	SetRegionId(v string) *ListEvaluationMetadataRequest
	GetRegionId() *string
	SetTopicCode(v string) *ListEvaluationMetadataRequest
	GetTopicCode() *string
}

type ListEvaluationMetadataRequest struct {
	// The language. The information is returned in the specified language. Valid values:
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LensCode *string `json:"LensCode,omitempty" xml:"LensCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ResourceUtilization
	TopicCode *string `json:"TopicCode,omitempty" xml:"TopicCode,omitempty"`
}

func (s ListEvaluationMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListEvaluationMetadataRequest) GetLensCode() *string {
	return s.LensCode
}

func (s *ListEvaluationMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEvaluationMetadataRequest) GetTopicCode() *string {
	return s.TopicCode
}

func (s *ListEvaluationMetadataRequest) SetLanguage(v string) *ListEvaluationMetadataRequest {
	s.Language = &v
	return s
}

func (s *ListEvaluationMetadataRequest) SetLensCode(v string) *ListEvaluationMetadataRequest {
	s.LensCode = &v
	return s
}

func (s *ListEvaluationMetadataRequest) SetRegionId(v string) *ListEvaluationMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationMetadataRequest) SetTopicCode(v string) *ListEvaluationMetadataRequest {
	s.TopicCode = &v
	return s
}

func (s *ListEvaluationMetadataRequest) Validate() error {
	return dara.Validate(s)
}
