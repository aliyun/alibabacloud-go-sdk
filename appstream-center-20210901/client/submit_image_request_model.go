// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionList(v []*string) *SubmitImageRequest
	GetDestinationRegionList() []*string
	SetTaskId(v string) *SubmitImageRequest
	GetTaskId() *string
}

type SubmitImageRequest struct {
	// The list of destination regions to which the image is copied.
	//
	// example:
	//
	// List
	DestinationRegionList []*string `json:"DestinationRegionList,omitempty" xml:"DestinationRegionList,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// tid-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitImageRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageRequest) GetDestinationRegionList() []*string {
	return s.DestinationRegionList
}

func (s *SubmitImageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitImageRequest) SetDestinationRegionList(v []*string) *SubmitImageRequest {
	s.DestinationRegionList = v
	return s
}

func (s *SubmitImageRequest) SetTaskId(v string) *SubmitImageRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitImageRequest) Validate() error {
	return dara.Validate(s)
}
