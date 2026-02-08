// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdList(v []*string) *GetSummaryInfoRequest
	GetInstanceIdList() []*string
}

type GetSummaryInfoRequest struct {
	// List of instance IDs
	//
	// example:
	//
	// []
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s GetSummaryInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryInfoRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *GetSummaryInfoRequest) SetInstanceIdList(v []*string) *GetSummaryInfoRequest {
	s.InstanceIdList = v
	return s
}

func (s *GetSummaryInfoRequest) Validate() error {
	return dara.Validate(s)
}
