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
	// The list of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to force a restart. You can force a restart if a cloud phone cannot be shut down due to system or network exceptions. Unsaved data may be lost.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// Specifies whether to skip parameter validation. A value of true indicates that validation is skipped. If the value is false or not specified, validation is performed.
	IgnoreParamValidation *bool `json:"IgnoreParamValidation,omitempty" xml:"IgnoreParamValidation,omitempty"`
	// **[Deprecated]*	- The sales mode.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
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
