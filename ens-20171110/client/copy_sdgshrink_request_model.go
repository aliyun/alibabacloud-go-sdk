// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIdsShrink(v string) *CopySDGShrinkRequest
	GetDestinationRegionIdsShrink() *string
	SetSDGId(v string) *CopySDGShrinkRequest
	GetSDGId() *string
}

type CopySDGShrinkRequest struct {
	// The destination nodes.
	//
	// This parameter is required.
	DestinationRegionIdsShrink *string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty"`
	// The ID of the SDG that you want to copy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s CopySDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *CopySDGShrinkRequest) GetDestinationRegionIdsShrink() *string {
	return s.DestinationRegionIdsShrink
}

func (s *CopySDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *CopySDGShrinkRequest) SetDestinationRegionIdsShrink(v string) *CopySDGShrinkRequest {
	s.DestinationRegionIdsShrink = &v
	return s
}

func (s *CopySDGShrinkRequest) SetSDGId(v string) *CopySDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *CopySDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
