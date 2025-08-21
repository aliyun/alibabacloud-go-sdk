// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeletePresetRequest
	GetId() *string
	SetOwnerId(v int64) *DeletePresetRequest
	GetOwnerId() *int64
	SetPresetId(v string) *DeletePresetRequest
	GetPresetId() *string
}

type DeletePresetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PresetId *string `json:"PresetId,omitempty" xml:"PresetId,omitempty"`
}

func (s DeletePresetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePresetRequest) GoString() string {
	return s.String()
}

func (s *DeletePresetRequest) GetId() *string {
	return s.Id
}

func (s *DeletePresetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePresetRequest) GetPresetId() *string {
	return s.PresetId
}

func (s *DeletePresetRequest) SetId(v string) *DeletePresetRequest {
	s.Id = &v
	return s
}

func (s *DeletePresetRequest) SetOwnerId(v int64) *DeletePresetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePresetRequest) SetPresetId(v string) *DeletePresetRequest {
	s.PresetId = &v
	return s
}

func (s *DeletePresetRequest) Validate() error {
	return dara.Validate(s)
}
