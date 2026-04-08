// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUpdateDesensStatusListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensStatus(v int32) *DsgUpdateDesensStatusListRequest
	GetDesensStatus() *int32
	SetIds(v []*int32) *DsgUpdateDesensStatusListRequest
	GetIds() []*int32
}

type DsgUpdateDesensStatusListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesensStatus *int32 `json:"DesensStatus,omitempty" xml:"DesensStatus,omitempty"`
	// This parameter is required.
	Ids []*int32 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DsgUpdateDesensStatusListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUpdateDesensStatusListRequest) GoString() string {
	return s.String()
}

func (s *DsgUpdateDesensStatusListRequest) GetDesensStatus() *int32 {
	return s.DesensStatus
}

func (s *DsgUpdateDesensStatusListRequest) GetIds() []*int32 {
	return s.Ids
}

func (s *DsgUpdateDesensStatusListRequest) SetDesensStatus(v int32) *DsgUpdateDesensStatusListRequest {
	s.DesensStatus = &v
	return s
}

func (s *DsgUpdateDesensStatusListRequest) SetIds(v []*int32) *DsgUpdateDesensStatusListRequest {
	s.Ids = v
	return s
}

func (s *DsgUpdateDesensStatusListRequest) Validate() error {
	return dara.Validate(s)
}
