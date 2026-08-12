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
	// The region ID of the access control instance. You can call the DescribeRegions operation to query the region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
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
	// Specifies whether the request is from the ECS console. Valid values:
	//
	// - **true**: The request is from the ECS console.
	//
	// - **false**: The request is not from the ECS console.
	//
	// example:
	//
	// true
	FromEcs *bool `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason for applying for the trial.
	RequestForm *CreateSasTrialRequestSdkRequestRequestForm `json:"RequestForm,omitempty" xml:"RequestForm,omitempty" type:"Struct"`
	// The trial type. Valid values:
	//
	// - **0**: trial not allowed
	//
	// - **1**: first trial
	//
	// - **2**: second trial
	//
	//
	// > Call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain this parameter. The trial can be started only when the value is not 0.
	//
	// example:
	//
	// 1
	TryType *int32 `json:"TryType,omitempty" xml:"TryType,omitempty"`
	// The trial edition. Valid values:
	//
	// - **3**: Enterprise Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// > Call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain this parameter.
	//
	// example:
	//
	// 3
	TryVersion *int32 `json:"TryVersion,omitempty" xml:"TryVersion,omitempty"`
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
	// The reason for applying for the trial.
	//
	// example:
	//
	// for poc
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
