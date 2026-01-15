// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAndroidInstancesInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *RebootAndroidInstancesInGroupRequest
	GetAndroidInstanceIds() []*string
	SetForceStop(v bool) *RebootAndroidInstancesInGroupRequest
	GetForceStop() *bool
	SetIgnoreParamValidation(v bool) *RebootAndroidInstancesInGroupRequest
	GetIgnoreParamValidation() *bool
	SetSaleMode(v string) *RebootAndroidInstancesInGroupRequest
	GetSaleMode() *string
}

type RebootAndroidInstancesInGroupRequest struct {
	// The IDs of the cloud phone instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enforce a restart operation. If a cloud phone instance fails to stop due to system or network issues, a forced restart can be triggered, though it may result in data loss.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ForceStop             *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	IgnoreParamValidation *bool   `json:"IgnoreParamValidation,omitempty" xml:"IgnoreParamValidation,omitempty"`
	SaleMode              *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s RebootAndroidInstancesInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupRequest) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *RebootAndroidInstancesInGroupRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RebootAndroidInstancesInGroupRequest) GetIgnoreParamValidation() *bool {
	return s.IgnoreParamValidation
}

func (s *RebootAndroidInstancesInGroupRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *RebootAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *RebootAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *RebootAndroidInstancesInGroupRequest) SetForceStop(v bool) *RebootAndroidInstancesInGroupRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootAndroidInstancesInGroupRequest) SetIgnoreParamValidation(v bool) *RebootAndroidInstancesInGroupRequest {
	s.IgnoreParamValidation = &v
	return s
}

func (s *RebootAndroidInstancesInGroupRequest) SetSaleMode(v string) *RebootAndroidInstancesInGroupRequest {
	s.SaleMode = &v
	return s
}

func (s *RebootAndroidInstancesInGroupRequest) Validate() error {
	return dara.Validate(s)
}
