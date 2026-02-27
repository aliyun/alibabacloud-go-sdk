// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceAppMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddDataServiceAppMemberShrinkRequest
	GetAddCommandShrink() *string
	SetOpTenantId(v int64) *AddDataServiceAppMemberShrinkRequest
	GetOpTenantId() *int64
}

type AddDataServiceAppMemberShrinkRequest struct {
	// This parameter is required.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddDataServiceAppMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddDataServiceAppMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddDataServiceAppMemberShrinkRequest) SetAddCommandShrink(v string) *AddDataServiceAppMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddDataServiceAppMemberShrinkRequest) SetOpTenantId(v int64) *AddDataServiceAppMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddDataServiceAppMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
