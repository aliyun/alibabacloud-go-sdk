// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumeOffsetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResetTime(v string) *ResetConsumeOffsetRequest
	GetResetTime() *string
	SetResetType(v string) *ResetConsumeOffsetRequest
	GetResetType() *string
}

type ResetConsumeOffsetRequest struct {
	// The time when the consumer offset is reset.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ResetTime *string `json:"resetTime,omitempty" xml:"resetTime,omitempty"`
	// The method that is used to reset the consumer offset. Valid values: LATEST_OFFSET and SPECIFIED_TIME.
	//
	// example:
	//
	// LATEST_OFFSET
	ResetType *string `json:"resetType,omitempty" xml:"resetType,omitempty"`
}

func (s ResetConsumeOffsetRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumeOffsetRequest) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetRequest) GetResetTime() *string {
	return s.ResetTime
}

func (s *ResetConsumeOffsetRequest) GetResetType() *string {
	return s.ResetType
}

func (s *ResetConsumeOffsetRequest) SetResetTime(v string) *ResetConsumeOffsetRequest {
	s.ResetTime = &v
	return s
}

func (s *ResetConsumeOffsetRequest) SetResetType(v string) *ResetConsumeOffsetRequest {
	s.ResetType = &v
	return s
}

func (s *ResetConsumeOffsetRequest) Validate() error {
	return dara.Validate(s)
}
