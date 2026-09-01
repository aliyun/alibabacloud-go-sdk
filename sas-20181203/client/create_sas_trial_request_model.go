// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSasTrialRequest
	GetClientToken() *string
	SetFromEcs(v bool) *CreateSasTrialRequest
	GetFromEcs() *bool
	SetLang(v string) *CreateSasTrialRequest
	GetLang() *string
	SetRequestForm(v *CreateSasTrialRequestRequestForm) *CreateSasTrialRequest
	GetRequestForm() *CreateSasTrialRequestRequestForm
	SetTryType(v int32) *CreateSasTrialRequest
	GetTryType() *int32
	SetTryVersion(v int32) *CreateSasTrialRequest
	GetTryVersion() *int32
}

type CreateSasTrialRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether the request is from the ECS console. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	FromEcs *bool `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason for applying for a trial. This parameter is required for a second trial.
	RequestForm *CreateSasTrialRequestRequestForm `json:"RequestForm,omitempty" xml:"RequestForm,omitempty" type:"Struct"`
	// The trial type. Valid values:
	//
	// - **0**: trial not allowed.
	//
	// - **1**: first trial.
	//
	// - **2**: second trial.
	//
	//
	// > Call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain this parameter. A trial can be started only when the value is not 0.
	//
	// example:
	//
	// 1
	TryType *int32 `json:"TryType,omitempty" xml:"TryType,omitempty"`
	// The trial version. Valid values:
	//
	// - **3**: Enterprise Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// >Call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain this parameter.
	//
	// example:
	//
	// 7
	TryVersion *int32 `json:"TryVersion,omitempty" xml:"TryVersion,omitempty"`
}

func (s CreateSasTrialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialRequest) GoString() string {
	return s.String()
}

func (s *CreateSasTrialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSasTrialRequest) GetFromEcs() *bool {
	return s.FromEcs
}

func (s *CreateSasTrialRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSasTrialRequest) GetRequestForm() *CreateSasTrialRequestRequestForm {
	return s.RequestForm
}

func (s *CreateSasTrialRequest) GetTryType() *int32 {
	return s.TryType
}

func (s *CreateSasTrialRequest) GetTryVersion() *int32 {
	return s.TryVersion
}

func (s *CreateSasTrialRequest) SetClientToken(v string) *CreateSasTrialRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSasTrialRequest) SetFromEcs(v bool) *CreateSasTrialRequest {
	s.FromEcs = &v
	return s
}

func (s *CreateSasTrialRequest) SetLang(v string) *CreateSasTrialRequest {
	s.Lang = &v
	return s
}

func (s *CreateSasTrialRequest) SetRequestForm(v *CreateSasTrialRequestRequestForm) *CreateSasTrialRequest {
	s.RequestForm = v
	return s
}

func (s *CreateSasTrialRequest) SetTryType(v int32) *CreateSasTrialRequest {
	s.TryType = &v
	return s
}

func (s *CreateSasTrialRequest) SetTryVersion(v int32) *CreateSasTrialRequest {
	s.TryVersion = &v
	return s
}

func (s *CreateSasTrialRequest) Validate() error {
	if s.RequestForm != nil {
		if err := s.RequestForm.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSasTrialRequestRequestForm struct {
	// The reason for applying for a trial.
	//
	// example:
	//
	// for poc
	TryReason *string `json:"TryReason,omitempty" xml:"TryReason,omitempty"`
}

func (s CreateSasTrialRequestRequestForm) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialRequestRequestForm) GoString() string {
	return s.String()
}

func (s *CreateSasTrialRequestRequestForm) GetTryReason() *string {
	return s.TryReason
}

func (s *CreateSasTrialRequestRequestForm) SetTryReason(v string) *CreateSasTrialRequestRequestForm {
	s.TryReason = &v
	return s
}

func (s *CreateSasTrialRequestRequestForm) Validate() error {
	return dara.Validate(s)
}
