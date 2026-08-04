// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOpResultDTO interface {
	dara.Model
	String() string
	GoString() string
	SetFailed(v []*BatchFailedItemDTO) *BatchOpResultDTO
	GetFailed() []*BatchFailedItemDTO
	SetSucceeded(v []*int64) *BatchOpResultDTO
	GetSucceeded() []*int64
}

type BatchOpResultDTO struct {
	// example:
	//
	// []
	Failed []*BatchFailedItemDTO `json:"failed,omitempty" xml:"failed,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Succeeded []*int64 `json:"succeeded,omitempty" xml:"succeeded,omitempty" type:"Repeated"`
}

func (s BatchOpResultDTO) String() string {
	return dara.Prettify(s)
}

func (s BatchOpResultDTO) GoString() string {
	return s.String()
}

func (s *BatchOpResultDTO) GetFailed() []*BatchFailedItemDTO {
	return s.Failed
}

func (s *BatchOpResultDTO) GetSucceeded() []*int64 {
	return s.Succeeded
}

func (s *BatchOpResultDTO) SetFailed(v []*BatchFailedItemDTO) *BatchOpResultDTO {
	s.Failed = v
	return s
}

func (s *BatchOpResultDTO) SetSucceeded(v []*int64) *BatchOpResultDTO {
	s.Succeeded = v
	return s
}

func (s *BatchOpResultDTO) Validate() error {
	if s.Failed != nil {
		for _, item := range s.Failed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
