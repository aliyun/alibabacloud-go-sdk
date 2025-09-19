// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// Specifies whether the request is redirected from the Elastic Compute Service (ECS) console. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FromEcs *bool `json:"FromEcs,omitempty" xml:"FromEcs,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why you apply for the trial. You must specify the reason for the second trial.
	RequestFormShrink *string `json:"RequestForm,omitempty" xml:"RequestForm,omitempty"`
	// The trial type. Valid values:
	//
	// 	- **0**: trial prohibited
	//
	// 	- **1**: first trial
	//
	// 	- **2**: second trial
	//
	// >  You can call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain the trial type. You can start a trial only if this parameter is not set to 0.
	//
	// example:
	//
	// 1
	TryType *int32 `json:"TryType,omitempty" xml:"TryType,omitempty"`
	// The trial edition. Valid values:
	//
	// 	- **3**: Enterprise
	//
	// 	- **7**: Ultimate
	//
	// >  You can call the [GetCanTrySas](https://help.aliyun.com/document_detail/2623574.html) operation to obtain the trial edition.
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
