// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetPresetRequest
	GetId() *string
	SetOwnerId(v int64) *SetPresetRequest
	GetOwnerId() *int64
	SetPresetId(v string) *SetPresetRequest
	GetPresetId() *string
}

type SetPresetRequest struct {
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

func (s SetPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPresetRequest) GoString() string {
	return s.String()
}

func (s *SetPresetRequest) GetId() *string {
	return s.Id
}

func (s *SetPresetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetPresetRequest) GetPresetId() *string {
	return s.PresetId
}

func (s *SetPresetRequest) SetId(v string) *SetPresetRequest {
	s.Id = &v
	return s
}

func (s *SetPresetRequest) SetOwnerId(v int64) *SetPresetRequest {
	s.OwnerId = &v
	return s
}

func (s *SetPresetRequest) SetPresetId(v string) *SetPresetRequest {
	s.PresetId = &v
	return s
}

func (s *SetPresetRequest) Validate() error {
	return dara.Validate(s)
}
