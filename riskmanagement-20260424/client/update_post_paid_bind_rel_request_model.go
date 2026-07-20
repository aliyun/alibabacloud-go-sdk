// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostPaidBindRelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdatePostPaidBindRelRequest
	GetRegionId() *string
	SetSdkRequest(v *UpdatePostPaidBindRelRequestSdkRequest) *UpdatePostPaidBindRelRequest
	GetSdkRequest() *UpdatePostPaidBindRelRequestSdkRequest
}

type UpdatePostPaidBindRelRequest struct {
	// example:
	//
	// cn-huhehaote
	RegionId   *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequest *UpdatePostPaidBindRelRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s UpdatePostPaidBindRelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePostPaidBindRelRequest) GetSdkRequest() *UpdatePostPaidBindRelRequestSdkRequest {
	return s.SdkRequest
}

func (s *UpdatePostPaidBindRelRequest) SetRegionId(v string) *UpdatePostPaidBindRelRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetSdkRequest(v *UpdatePostPaidBindRelRequestSdkRequest) *UpdatePostPaidBindRelRequest {
	s.SdkRequest = v
	return s
}

func (s *UpdatePostPaidBindRelRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePostPaidBindRelRequestSdkRequest struct {
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// example:
	//
	// 3
	AutoBindVersion *int32                                              `json:"AutoBindVersion,omitempty" xml:"AutoBindVersion,omitempty"`
	BindAction      []*UpdatePostPaidBindRelRequestSdkRequestBindAction `json:"BindAction,omitempty" xml:"BindAction,omitempty" type:"Repeated"`
	// example:
	//
	// false
	UpdateIfNecessary *bool `json:"UpdateIfNecessary,omitempty" xml:"UpdateIfNecessary,omitempty"`
}

func (s UpdatePostPaidBindRelRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) GetAutoBindVersion() *int32 {
	return s.AutoBindVersion
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) GetBindAction() []*UpdatePostPaidBindRelRequestSdkRequestBindAction {
	return s.BindAction
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) GetUpdateIfNecessary() *bool {
	return s.UpdateIfNecessary
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) SetAutoBind(v int32) *UpdatePostPaidBindRelRequestSdkRequest {
	s.AutoBind = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) SetAutoBindVersion(v int32) *UpdatePostPaidBindRelRequestSdkRequest {
	s.AutoBindVersion = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) SetBindAction(v []*UpdatePostPaidBindRelRequestSdkRequestBindAction) *UpdatePostPaidBindRelRequestSdkRequest {
	s.BindAction = v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) SetUpdateIfNecessary(v bool) *UpdatePostPaidBindRelRequestSdkRequest {
	s.UpdateIfNecessary = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequest) Validate() error {
	if s.BindAction != nil {
		for _, item := range s.BindAction {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePostPaidBindRelRequestSdkRequestBindAction struct {
	// example:
	//
	// true
	BindAll  *bool     `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdatePostPaidBindRelRequestSdkRequestBindAction) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelRequestSdkRequestBindAction) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) GetBindAll() *bool {
	return s.BindAll
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) GetUuidList() []*string {
	return s.UuidList
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) GetVersion() *string {
	return s.Version
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) SetBindAll(v bool) *UpdatePostPaidBindRelRequestSdkRequestBindAction {
	s.BindAll = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) SetUuidList(v []*string) *UpdatePostPaidBindRelRequestSdkRequestBindAction {
	s.UuidList = v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) SetVersion(v string) *UpdatePostPaidBindRelRequestSdkRequestBindAction {
	s.Version = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestSdkRequestBindAction) Validate() error {
	return dara.Validate(s)
}
