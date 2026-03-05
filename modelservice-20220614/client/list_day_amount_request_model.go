// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDayAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListDayAmountRequest
	GetEndTime() *string
	SetSceneType(v string) *ListDayAmountRequest
	GetSceneType() *string
	SetStartTime(v string) *ListDayAmountRequest
	GetStartTime() *string
}

type ListDayAmountRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// sales_pick
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDayAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDayAmountRequest) GoString() string {
	return s.String()
}

func (s *ListDayAmountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDayAmountRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *ListDayAmountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDayAmountRequest) SetEndTime(v string) *ListDayAmountRequest {
	s.EndTime = &v
	return s
}

func (s *ListDayAmountRequest) SetSceneType(v string) *ListDayAmountRequest {
	s.SceneType = &v
	return s
}

func (s *ListDayAmountRequest) SetStartTime(v string) *ListDayAmountRequest {
	s.StartTime = &v
	return s
}

func (s *ListDayAmountRequest) Validate() error {
	return dara.Validate(s)
}
