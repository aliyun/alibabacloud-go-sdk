// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGotoPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GotoPresetRequest
	GetId() *string
	SetOwnerId(v int64) *GotoPresetRequest
	GetOwnerId() *int64
	SetPresetId(v string) *GotoPresetRequest
	GetPresetId() *string
}

type GotoPresetRequest struct {
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

func (s GotoPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s GotoPresetRequest) GoString() string {
	return s.String()
}

func (s *GotoPresetRequest) GetId() *string {
	return s.Id
}

func (s *GotoPresetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GotoPresetRequest) GetPresetId() *string {
	return s.PresetId
}

func (s *GotoPresetRequest) SetId(v string) *GotoPresetRequest {
	s.Id = &v
	return s
}

func (s *GotoPresetRequest) SetOwnerId(v int64) *GotoPresetRequest {
	s.OwnerId = &v
	return s
}

func (s *GotoPresetRequest) SetPresetId(v string) *GotoPresetRequest {
	s.PresetId = &v
	return s
}

func (s *GotoPresetRequest) Validate() error {
	return dara.Validate(s)
}
