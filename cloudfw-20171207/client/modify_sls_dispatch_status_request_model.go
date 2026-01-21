// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchValue(v string) *ModifySlsDispatchStatusRequest
	GetDispatchValue() *string
	SetEnableStatus(v bool) *ModifySlsDispatchStatusRequest
	GetEnableStatus() *bool
	SetFilterKeys(v string) *ModifySlsDispatchStatusRequest
	GetFilterKeys() *string
	SetNewRegionId(v string) *ModifySlsDispatchStatusRequest
	GetNewRegionId() *string
	SetSite(v string) *ModifySlsDispatchStatusRequest
	GetSite() *string
}

type ModifySlsDispatchStatusRequest struct {
	// example:
	//
	// internet_log
	DispatchValue *string `json:"DispatchValue,omitempty" xml:"DispatchValue,omitempty"`
	// example:
	//
	// true
	EnableStatus *bool `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// attack,acl
	FilterKeys *string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty"`
	// example:
	//
	// cn-shanghai
	NewRegionId *string `json:"NewRegionId,omitempty" xml:"NewRegionId,omitempty"`
	// example:
	//
	// cn
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s ModifySlsDispatchStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchStatusRequest) GetDispatchValue() *string {
	return s.DispatchValue
}

func (s *ModifySlsDispatchStatusRequest) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *ModifySlsDispatchStatusRequest) GetFilterKeys() *string {
	return s.FilterKeys
}

func (s *ModifySlsDispatchStatusRequest) GetNewRegionId() *string {
	return s.NewRegionId
}

func (s *ModifySlsDispatchStatusRequest) GetSite() *string {
	return s.Site
}

func (s *ModifySlsDispatchStatusRequest) SetDispatchValue(v string) *ModifySlsDispatchStatusRequest {
	s.DispatchValue = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetEnableStatus(v bool) *ModifySlsDispatchStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetFilterKeys(v string) *ModifySlsDispatchStatusRequest {
	s.FilterKeys = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetNewRegionId(v string) *ModifySlsDispatchStatusRequest {
	s.NewRegionId = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) SetSite(v string) *ModifySlsDispatchStatusRequest {
	s.Site = &v
	return s
}

func (s *ModifySlsDispatchStatusRequest) Validate() error {
	return dara.Validate(s)
}
