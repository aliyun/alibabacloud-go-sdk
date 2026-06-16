// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BindAuthToMachineRequest
	GetRegionId() *string
	SetSdkRequest(v *BindAuthToMachineRequestSdkRequest) *BindAuthToMachineRequest
	GetSdkRequest() *BindAuthToMachineRequestSdkRequest
}

type BindAuthToMachineRequest struct {
	RegionId   *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequest *BindAuthToMachineRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s BindAuthToMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindAuthToMachineRequest) GetSdkRequest() *BindAuthToMachineRequestSdkRequest {
	return s.SdkRequest
}

func (s *BindAuthToMachineRequest) SetRegionId(v string) *BindAuthToMachineRequest {
	s.RegionId = &v
	return s
}

func (s *BindAuthToMachineRequest) SetSdkRequest(v *BindAuthToMachineRequestSdkRequest) *BindAuthToMachineRequest {
	s.SdkRequest = v
	return s
}

func (s *BindAuthToMachineRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAuthToMachineRequestSdkRequest struct {
	AuthVersion    *int32    `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	AutoBind       *int32    `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	Bind           []*string `json:"Bind,omitempty" xml:"Bind,omitempty" type:"Repeated"`
	BindAll        *bool     `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	Criteria       *string   `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	IsPreBind      *int32    `json:"IsPreBind,omitempty" xml:"IsPreBind,omitempty"`
	LogicalExp     *string   `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	NtmVersion     *int64    `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	PreBindOrderId *int64    `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	UnBind         []*string `json:"UnBind,omitempty" xml:"UnBind,omitempty" type:"Repeated"`
}

func (s BindAuthToMachineRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineRequestSdkRequest) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *BindAuthToMachineRequestSdkRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *BindAuthToMachineRequestSdkRequest) GetBind() []*string {
	return s.Bind
}

func (s *BindAuthToMachineRequestSdkRequest) GetBindAll() *bool {
	return s.BindAll
}

func (s *BindAuthToMachineRequestSdkRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *BindAuthToMachineRequestSdkRequest) GetIsPreBind() *int32 {
	return s.IsPreBind
}

func (s *BindAuthToMachineRequestSdkRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *BindAuthToMachineRequestSdkRequest) GetNtmVersion() *int64 {
	return s.NtmVersion
}

func (s *BindAuthToMachineRequestSdkRequest) GetPreBindOrderId() *int64 {
	return s.PreBindOrderId
}

func (s *BindAuthToMachineRequestSdkRequest) GetUnBind() []*string {
	return s.UnBind
}

func (s *BindAuthToMachineRequestSdkRequest) SetAuthVersion(v int32) *BindAuthToMachineRequestSdkRequest {
	s.AuthVersion = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetAutoBind(v int32) *BindAuthToMachineRequestSdkRequest {
	s.AutoBind = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetBind(v []*string) *BindAuthToMachineRequestSdkRequest {
	s.Bind = v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetBindAll(v bool) *BindAuthToMachineRequestSdkRequest {
	s.BindAll = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetCriteria(v string) *BindAuthToMachineRequestSdkRequest {
	s.Criteria = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetIsPreBind(v int32) *BindAuthToMachineRequestSdkRequest {
	s.IsPreBind = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetLogicalExp(v string) *BindAuthToMachineRequestSdkRequest {
	s.LogicalExp = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetNtmVersion(v int64) *BindAuthToMachineRequestSdkRequest {
	s.NtmVersion = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetPreBindOrderId(v int64) *BindAuthToMachineRequestSdkRequest {
	s.PreBindOrderId = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetUnBind(v []*string) *BindAuthToMachineRequestSdkRequest {
	s.UnBind = v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
