// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DeleteForwardEntryRequest
	GetForwardEntryId() *string
	SetForwardTableId(v string) *DeleteForwardEntryRequest
	GetForwardTableId() *string
	SetRegionId(v string) *DeleteForwardEntryRequest
	GetRegionId() *string
}

type DeleteForwardEntryRequest struct {
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DeleteForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DeleteForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteForwardEntryRequest) SetForwardEntryId(v string) *DeleteForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetForwardTableId(v string) *DeleteForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetRegionId(v string) *DeleteForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
