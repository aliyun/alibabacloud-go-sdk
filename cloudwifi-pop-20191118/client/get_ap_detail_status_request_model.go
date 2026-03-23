// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetApDetailStatusRequest
	GetAppCode() *string
	SetAppName(v string) *GetApDetailStatusRequest
	GetAppName() *string
	SetMac(v string) *GetApDetailStatusRequest
	GetMac() *string
	SetNeedApgroupInfo(v bool) *GetApDetailStatusRequest
	GetNeedApgroupInfo() *bool
	SetNeedRadioStatus(v bool) *GetApDetailStatusRequest
	GetNeedRadioStatus() *bool
}

type GetApDetailStatusRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// This parameter is required.
	NeedApgroupInfo *bool `json:"NeedApgroupInfo,omitempty" xml:"NeedApgroupInfo,omitempty"`
	// This parameter is required.
	NeedRadioStatus *bool `json:"NeedRadioStatus,omitempty" xml:"NeedRadioStatus,omitempty"`
}

func (s GetApDetailStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailStatusRequest) GoString() string {
	return s.String()
}

func (s *GetApDetailStatusRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApDetailStatusRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApDetailStatusRequest) GetMac() *string {
	return s.Mac
}

func (s *GetApDetailStatusRequest) GetNeedApgroupInfo() *bool {
	return s.NeedApgroupInfo
}

func (s *GetApDetailStatusRequest) GetNeedRadioStatus() *bool {
	return s.NeedRadioStatus
}

func (s *GetApDetailStatusRequest) SetAppCode(v string) *GetApDetailStatusRequest {
	s.AppCode = &v
	return s
}

func (s *GetApDetailStatusRequest) SetAppName(v string) *GetApDetailStatusRequest {
	s.AppName = &v
	return s
}

func (s *GetApDetailStatusRequest) SetMac(v string) *GetApDetailStatusRequest {
	s.Mac = &v
	return s
}

func (s *GetApDetailStatusRequest) SetNeedApgroupInfo(v bool) *GetApDetailStatusRequest {
	s.NeedApgroupInfo = &v
	return s
}

func (s *GetApDetailStatusRequest) SetNeedRadioStatus(v bool) *GetApDetailStatusRequest {
	s.NeedRadioStatus = &v
	return s
}

func (s *GetApDetailStatusRequest) Validate() error {
	return dara.Validate(s)
}
