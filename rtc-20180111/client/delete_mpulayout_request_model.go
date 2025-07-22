// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMPULayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMPULayoutRequest
	GetAppId() *string
	SetLayoutId(v int64) *DeleteMPULayoutRequest
	GetLayoutId() *int64
	SetOwnerId(v int64) *DeleteMPULayoutRequest
	GetOwnerId() *int64
}

type DeleteMPULayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteMPULayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMPULayoutRequest) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *DeleteMPULayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMPULayoutRequest) SetAppId(v string) *DeleteMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetLayoutId(v int64) *DeleteMPULayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetOwnerId(v int64) *DeleteMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMPULayoutRequest) Validate() error {
	return dara.Validate(s)
}
