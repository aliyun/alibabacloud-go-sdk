// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSasTrialShrinkRequest
	GetClientToken() *string
	SetFromEcs(v bool) *CreateSasTrialShrinkRequest
	GetFromEcs() *bool
	SetLang(v string) *CreateSasTrialShrinkRequest
	GetLang() *string
	SetRequestFormShrink(v string) *CreateSasTrialShrinkRequest
	GetRequestFormShrink() *string
	SetTryType(v int32) *CreateSasTrialShrinkRequest
	GetTryType() *int32
	SetTryVersion(v int32) *CreateSasTrialShrinkRequest
	GetTryVersion() *int32
}

type CreateSasTrialShrinkRequest struct {
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
	RequestFormShrink *string `json:"RequestForm,omitempty" xml:"RequestForm,omitempty"`
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

func (s CreateSasTrialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSasTrialShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSasTrialShrinkRequest) GetFromEcs() *bool {
	return s.FromEcs
}

func (s *CreateSasTrialShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSasTrialShrinkRequest) GetRequestFormShrink() *string {
	return s.RequestFormShrink
}

func (s *CreateSasTrialShrinkRequest) GetTryType() *int32 {
	return s.TryType
}

func (s *CreateSasTrialShrinkRequest) GetTryVersion() *int32 {
	return s.TryVersion
}

func (s *CreateSasTrialShrinkRequest) SetClientToken(v string) *CreateSasTrialShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetFromEcs(v bool) *CreateSasTrialShrinkRequest {
	s.FromEcs = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetLang(v string) *CreateSasTrialShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetRequestFormShrink(v string) *CreateSasTrialShrinkRequest {
	s.RequestFormShrink = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetTryType(v int32) *CreateSasTrialShrinkRequest {
	s.TryType = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetTryVersion(v int32) *CreateSasTrialShrinkRequest {
	s.TryVersion = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
