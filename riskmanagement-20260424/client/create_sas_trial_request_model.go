// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateSasTrialRequest
	GetRegionId() *string
	SetSdkRequest(v *CreateSasTrialRequestSdkRequest) *CreateSasTrialRequest
	GetSdkRequest() *CreateSasTrialRequestSdkRequest
}

type CreateSasTrialRequest struct {
	RegionId   *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequest *CreateSasTrialRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s CreateSasTrialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialRequest) GoString() string {
	return s.String()
}

func (s *CreateSasTrialRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSasTrialRequest) GetSdkRequest() *CreateSasTrialRequestSdkRequest {
	return s.SdkRequest
}

func (s *CreateSasTrialRequest) SetRegionId(v string) *CreateSasTrialRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSasTrialRequest) SetSdkRequest(v *CreateSasTrialRequestSdkRequest) *CreateSasTrialRequest {
	s.SdkRequest = v
	return s
}

func (s *CreateSasTrialRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSasTrialRequestSdkRequest struct {
	FromEcs     *bool                                       `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	Lang        *string                                     `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RequestForm *CreateSasTrialRequestSdkRequestRequestForm `json:"RequestForm,omitempty" xml:"RequestForm,omitempty" type:"Struct"`
	TryType     *int32                                      `json:"TryType,omitempty" xml:"TryType,omitempty"`
	TryVersion  *int32                                      `json:"TryVersion,omitempty" xml:"TryVersion,omitempty"`
}

func (s CreateSasTrialRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *CreateSasTrialRequestSdkRequest) GetFromEcs() *bool {
	return s.FromEcs
}

func (s *CreateSasTrialRequestSdkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSasTrialRequestSdkRequest) GetRequestForm() *CreateSasTrialRequestSdkRequestRequestForm {
	return s.RequestForm
}

func (s *CreateSasTrialRequestSdkRequest) GetTryType() *int32 {
	return s.TryType
}

func (s *CreateSasTrialRequestSdkRequest) GetTryVersion() *int32 {
	return s.TryVersion
}

func (s *CreateSasTrialRequestSdkRequest) SetFromEcs(v bool) *CreateSasTrialRequestSdkRequest {
	s.FromEcs = &v
	return s
}

func (s *CreateSasTrialRequestSdkRequest) SetLang(v string) *CreateSasTrialRequestSdkRequest {
	s.Lang = &v
	return s
}

func (s *CreateSasTrialRequestSdkRequest) SetRequestForm(v *CreateSasTrialRequestSdkRequestRequestForm) *CreateSasTrialRequestSdkRequest {
	s.RequestForm = v
	return s
}

func (s *CreateSasTrialRequestSdkRequest) SetTryType(v int32) *CreateSasTrialRequestSdkRequest {
	s.TryType = &v
	return s
}

func (s *CreateSasTrialRequestSdkRequest) SetTryVersion(v int32) *CreateSasTrialRequestSdkRequest {
	s.TryVersion = &v
	return s
}

func (s *CreateSasTrialRequestSdkRequest) Validate() error {
	if s.RequestForm != nil {
		if err := s.RequestForm.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSasTrialRequestSdkRequestRequestForm struct {
	TryReason *string `json:"TryReason,omitempty" xml:"TryReason,omitempty"`
}

func (s CreateSasTrialRequestSdkRequestRequestForm) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialRequestSdkRequestRequestForm) GoString() string {
	return s.String()
}

func (s *CreateSasTrialRequestSdkRequestRequestForm) GetTryReason() *string {
	return s.TryReason
}

func (s *CreateSasTrialRequestSdkRequestRequestForm) SetTryReason(v string) *CreateSasTrialRequestSdkRequestRequestForm {
	s.TryReason = &v
	return s
}

func (s *CreateSasTrialRequestSdkRequestRequestForm) Validate() error {
	return dara.Validate(s)
}
