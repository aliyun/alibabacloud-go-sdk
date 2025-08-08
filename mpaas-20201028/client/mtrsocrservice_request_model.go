// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMTRSOCRServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MTRSOCRServiceRequest
	GetAppId() *string
	SetImageRaw(v string) *MTRSOCRServiceRequest
	GetImageRaw() *string
	SetMask(v bool) *MTRSOCRServiceRequest
	GetMask() *bool
	SetTenantId(v string) *MTRSOCRServiceRequest
	GetTenantId() *string
	SetType(v string) *MTRSOCRServiceRequest
	GetType() *string
	SetWorkspaceId(v string) *MTRSOCRServiceRequest
	GetWorkspaceId() *string
}

type MTRSOCRServiceRequest struct {
	// example:
	//
	// ONEX8C7E7FA161089
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xydasf==ac
	ImageRaw *string `json:"ImageRaw,omitempty" xml:"ImageRaw,omitempty"`
	Mask     *bool   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// example:
	//
	// tabcaa
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID_CARD_OCR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MTRSOCRServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s MTRSOCRServiceRequest) GoString() string {
	return s.String()
}

func (s *MTRSOCRServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *MTRSOCRServiceRequest) GetImageRaw() *string {
	return s.ImageRaw
}

func (s *MTRSOCRServiceRequest) GetMask() *bool {
	return s.Mask
}

func (s *MTRSOCRServiceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MTRSOCRServiceRequest) GetType() *string {
	return s.Type
}

func (s *MTRSOCRServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MTRSOCRServiceRequest) SetAppId(v string) *MTRSOCRServiceRequest {
	s.AppId = &v
	return s
}

func (s *MTRSOCRServiceRequest) SetImageRaw(v string) *MTRSOCRServiceRequest {
	s.ImageRaw = &v
	return s
}

func (s *MTRSOCRServiceRequest) SetMask(v bool) *MTRSOCRServiceRequest {
	s.Mask = &v
	return s
}

func (s *MTRSOCRServiceRequest) SetTenantId(v string) *MTRSOCRServiceRequest {
	s.TenantId = &v
	return s
}

func (s *MTRSOCRServiceRequest) SetType(v string) *MTRSOCRServiceRequest {
	s.Type = &v
	return s
}

func (s *MTRSOCRServiceRequest) SetWorkspaceId(v string) *MTRSOCRServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MTRSOCRServiceRequest) Validate() error {
	return dara.Validate(s)
}
