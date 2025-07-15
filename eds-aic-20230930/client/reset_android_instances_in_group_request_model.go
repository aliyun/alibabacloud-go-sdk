// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAndroidInstancesInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest
	GetAndroidInstanceIds() []*string
	SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest
	GetSaleMode() *string
	SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest
	GetSettingResetType() *int32
}

type ResetAndroidInstancesInGroupRequest struct {
	// The IDs of the cloud phone instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	SaleMode           *string   `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	SettingResetType   *int32    `json:"SettingResetType,omitempty" xml:"SettingResetType,omitempty"`
}

func (s ResetAndroidInstancesInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupRequest) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *ResetAndroidInstancesInGroupRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *ResetAndroidInstancesInGroupRequest) GetSettingResetType() *int32 {
	return s.SettingResetType
}

func (s *ResetAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest {
	s.SaleMode = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest {
	s.SettingResetType = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) Validate() error {
	return dara.Validate(s)
}
