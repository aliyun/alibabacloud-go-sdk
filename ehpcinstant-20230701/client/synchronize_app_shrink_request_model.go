// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SynchronizeAppShrinkRequest
	GetAppId() *string
	SetTargetRegionIdsShrink(v string) *SynchronizeAppShrinkRequest
	GetTargetRegionIdsShrink() *string
}

type SynchronizeAppShrinkRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ci-vm-rYfypJKwlN9Y
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The list of arrays that are synchronized to the specified region. If \\"all\\" is included, it is synchronized to all other unsynchronized regions by default.
	TargetRegionIdsShrink *string `json:"TargetRegionIds,omitempty" xml:"TargetRegionIds,omitempty"`
}

func (s SynchronizeAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *SynchronizeAppShrinkRequest) GetTargetRegionIdsShrink() *string {
	return s.TargetRegionIdsShrink
}

func (s *SynchronizeAppShrinkRequest) SetAppId(v string) *SynchronizeAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SynchronizeAppShrinkRequest) SetTargetRegionIdsShrink(v string) *SynchronizeAppShrinkRequest {
	s.TargetRegionIdsShrink = &v
	return s
}

func (s *SynchronizeAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
