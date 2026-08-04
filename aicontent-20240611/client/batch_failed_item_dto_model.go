// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFailedItemDTO interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *BatchFailedItemDTO
	GetReason() *string
	SetUserId(v int64) *BatchFailedItemDTO
	GetUserId() *int64
}

type BatchFailedItemDTO struct {
	// example:
	//
	// Member node is missing
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchFailedItemDTO) String() string {
	return dara.Prettify(s)
}

func (s BatchFailedItemDTO) GoString() string {
	return s.String()
}

func (s *BatchFailedItemDTO) GetReason() *string {
	return s.Reason
}

func (s *BatchFailedItemDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *BatchFailedItemDTO) SetReason(v string) *BatchFailedItemDTO {
	s.Reason = &v
	return s
}

func (s *BatchFailedItemDTO) SetUserId(v int64) *BatchFailedItemDTO {
	s.UserId = &v
	return s
}

func (s *BatchFailedItemDTO) Validate() error {
	return dara.Validate(s)
}
