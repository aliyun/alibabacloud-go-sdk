// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SynchronizeAppRequest
	GetAppId() *string
	SetTargetRegionIds(v []*string) *SynchronizeAppRequest
	GetTargetRegionIds() []*string
}

type SynchronizeAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ci-vm-rYfypJKwlN9Y
	AppId           *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TargetRegionIds []*string `json:"TargetRegionIds,omitempty" xml:"TargetRegionIds,omitempty" type:"Repeated"`
}

func (s SynchronizeAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeAppRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *SynchronizeAppRequest) GetTargetRegionIds() []*string {
	return s.TargetRegionIds
}

func (s *SynchronizeAppRequest) SetAppId(v string) *SynchronizeAppRequest {
	s.AppId = &v
	return s
}

func (s *SynchronizeAppRequest) SetTargetRegionIds(v []*string) *SynchronizeAppRequest {
	s.TargetRegionIds = v
	return s
}

func (s *SynchronizeAppRequest) Validate() error {
	return dara.Validate(s)
}
