// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostPaidBindRelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBind(v int32) *UpdatePostPaidBindRelRequest
	GetAutoBind() *int32
	SetAutoBindVersion(v int32) *UpdatePostPaidBindRelRequest
	GetAutoBindVersion() *int32
	SetBindAction(v []*UpdatePostPaidBindRelRequestBindAction) *UpdatePostPaidBindRelRequest
	GetBindAction() []*UpdatePostPaidBindRelRequestBindAction
	SetUpdateIfNecessary(v bool) *UpdatePostPaidBindRelRequest
	GetUpdateIfNecessary() *bool
}

type UpdatePostPaidBindRelRequest struct {
	// Enable automatic binding for new assets. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// Version to automatically bind when adding new assets. Values:
	//
	// - **1**: Basic Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Antivirus Edition
	//
	// - **7**: Container Edition
	//
	// example:
	//
	// 3
	AutoBindVersion *int32 `json:"AutoBindVersion,omitempty" xml:"AutoBindVersion,omitempty"`
	// Parameters for the binding action.
	BindAction        []*UpdatePostPaidBindRelRequestBindAction `json:"BindAction,omitempty" xml:"BindAction,omitempty" type:"Repeated"`
	UpdateIfNecessary *bool                                     `json:"UpdateIfNecessary,omitempty" xml:"UpdateIfNecessary,omitempty"`
}

func (s UpdatePostPaidBindRelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *UpdatePostPaidBindRelRequest) GetAutoBindVersion() *int32 {
	return s.AutoBindVersion
}

func (s *UpdatePostPaidBindRelRequest) GetBindAction() []*UpdatePostPaidBindRelRequestBindAction {
	return s.BindAction
}

func (s *UpdatePostPaidBindRelRequest) GetUpdateIfNecessary() *bool {
	return s.UpdateIfNecessary
}

func (s *UpdatePostPaidBindRelRequest) SetAutoBind(v int32) *UpdatePostPaidBindRelRequest {
	s.AutoBind = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetAutoBindVersion(v int32) *UpdatePostPaidBindRelRequest {
	s.AutoBindVersion = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetBindAction(v []*UpdatePostPaidBindRelRequestBindAction) *UpdatePostPaidBindRelRequest {
	s.BindAction = v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetUpdateIfNecessary(v bool) *UpdatePostPaidBindRelRequest {
	s.UpdateIfNecessary = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) Validate() error {
	return dara.Validate(s)
}

type UpdatePostPaidBindRelRequestBindAction struct {
	// Whether to bind all. Default is **false**. Values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// List of specified server UUIDs.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
	// The Cloud Security Center protection version that needs to be bound. Values:
	//
	// - **1**: Basic Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Antivirus Edition
	//
	// - **7**: Container Edition
	//
	// example:
	//
	// 3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdatePostPaidBindRelRequestBindAction) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelRequestBindAction) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelRequestBindAction) GetBindAll() *bool {
	return s.BindAll
}

func (s *UpdatePostPaidBindRelRequestBindAction) GetUuidList() []*string {
	return s.UuidList
}

func (s *UpdatePostPaidBindRelRequestBindAction) GetVersion() *string {
	return s.Version
}

func (s *UpdatePostPaidBindRelRequestBindAction) SetBindAll(v bool) *UpdatePostPaidBindRelRequestBindAction {
	s.BindAll = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestBindAction) SetUuidList(v []*string) *UpdatePostPaidBindRelRequestBindAction {
	s.UuidList = v
	return s
}

func (s *UpdatePostPaidBindRelRequestBindAction) SetVersion(v string) *UpdatePostPaidBindRelRequestBindAction {
	s.Version = &v
	return s
}

func (s *UpdatePostPaidBindRelRequestBindAction) Validate() error {
	return dara.Validate(s)
}
