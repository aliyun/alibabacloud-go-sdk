// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCallbackRequest
	GetId() *int64
	SetRegionId(v string) *DeleteCallbackRequest
	GetRegionId() *string
}

type DeleteCallbackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1480
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeleteCallbackRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCallbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCallbackRequest) SetId(v int64) *DeleteCallbackRequest {
	s.Id = &v
	return s
}

func (s *DeleteCallbackRequest) SetRegionId(v string) *DeleteCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCallbackRequest) Validate() error {
	return dara.Validate(s)
}
