// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationRegionIds(v []*string) *CopySDGRequest
	GetDestinationRegionIds() []*string
	SetSDGId(v string) *CopySDGRequest
	GetSDGId() *string
}

type CopySDGRequest struct {
	// The destination nodes.
	//
	// This parameter is required.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The ID of the SDG that you want to copy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s CopySDGRequest) String() string {
	return dara.Prettify(s)
}

func (s CopySDGRequest) GoString() string {
	return s.String()
}

func (s *CopySDGRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *CopySDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *CopySDGRequest) SetDestinationRegionIds(v []*string) *CopySDGRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *CopySDGRequest) SetSDGId(v string) *CopySDGRequest {
	s.SDGId = &v
	return s
}

func (s *CopySDGRequest) Validate() error {
	return dara.Validate(s)
}
