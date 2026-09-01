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
	SetClientToken(v string) *UpdatePostPaidBindRelRequest
	GetClientToken() *string
	SetProductCode(v string) *UpdatePostPaidBindRelRequest
	GetProductCode() *string
	SetUpdateIfNecessary(v bool) *UpdatePostPaidBindRelRequest
	GetUpdateIfNecessary() *bool
}

type UpdatePostPaidBindRelRequest struct {
	// Specifies whether to enable automatic binding for new assets. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// The edition to automatically bind when new assets are added. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
	//
	// example:
	//
	// 3
	AutoBindVersion *int32 `json:"AutoBindVersion,omitempty" xml:"AutoBindVersion,omitempty"`
	// The binding action parameter.
	BindAction []*UpdatePostPaidBindRelRequestBindAction `json:"BindAction,omitempty" xml:"BindAction,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to forcibly upgrade the edition.
	//
	// example:
	//
	// false
	UpdateIfNecessary *bool `json:"UpdateIfNecessary,omitempty" xml:"UpdateIfNecessary,omitempty"`
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

func (s *UpdatePostPaidBindRelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePostPaidBindRelRequest) GetProductCode() *string {
	return s.ProductCode
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

func (s *UpdatePostPaidBindRelRequest) SetClientToken(v string) *UpdatePostPaidBindRelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetProductCode(v string) *UpdatePostPaidBindRelRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) SetUpdateIfNecessary(v bool) *UpdatePostPaidBindRelRequest {
	s.UpdateIfNecessary = &v
	return s
}

func (s *UpdatePostPaidBindRelRequest) Validate() error {
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

type UpdatePostPaidBindRelRequestBindAction struct {
	// Specifies whether to bind all servers. Default value: **false**. Valid values:
	//
	// - **true**: yes
	//
	// - **false**: no
	//
	// example:
	//
	// true
	BindAll  *bool   `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	FreeType *string `json:"FreeType,omitempty" xml:"FreeType,omitempty"`
	// The list of server UUIDs.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
	// The protection edition of Security Center to bind. Valid values:
	//
	// - **1**: Free Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **5**: Advanced Edition
	//
	// - **6**: Anti-virus Edition
	//
	// - **7**: Ultimate Edition
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

func (s *UpdatePostPaidBindRelRequestBindAction) GetFreeType() *string {
	return s.FreeType
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

func (s *UpdatePostPaidBindRelRequestBindAction) SetFreeType(v string) *UpdatePostPaidBindRelRequestBindAction {
	s.FreeType = &v
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
