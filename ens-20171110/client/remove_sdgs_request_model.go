// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *RemoveSDGsRequest
	GetInstanceIds() []*string
	SetSdgIds(v []*string) *RemoveSDGsRequest
	GetSdgIds() []*string
}

type RemoveSDGsRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	SdgIds      []*string `json:"SdgIds,omitempty" xml:"SdgIds,omitempty" type:"Repeated"`
}

func (s RemoveSDGsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsRequest) GoString() string {
	return s.String()
}

func (s *RemoveSDGsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveSDGsRequest) GetSdgIds() []*string {
	return s.SdgIds
}

func (s *RemoveSDGsRequest) SetInstanceIds(v []*string) *RemoveSDGsRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveSDGsRequest) SetSdgIds(v []*string) *RemoveSDGsRequest {
	s.SdgIds = v
	return s
}

func (s *RemoveSDGsRequest) Validate() error {
	return dara.Validate(s)
}
