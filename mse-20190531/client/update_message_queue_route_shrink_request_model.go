// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageQueueRouteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateMessageQueueRouteShrinkRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateMessageQueueRouteShrinkRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMessageQueueRouteShrinkRequest
	GetAppName() *string
	SetEnable(v bool) *UpdateMessageQueueRouteShrinkRequest
	GetEnable() *bool
	SetFilterSide(v string) *UpdateMessageQueueRouteShrinkRequest
	GetFilterSide() *string
	SetGrayBaseTagsShrink(v string) *UpdateMessageQueueRouteShrinkRequest
	GetGrayBaseTagsShrink() *string
	SetNamespace(v string) *UpdateMessageQueueRouteShrinkRequest
	GetNamespace() *string
	SetRegion(v string) *UpdateMessageQueueRouteShrinkRequest
	GetRegion() *string
	SetTagsShrink(v string) *UpdateMessageQueueRouteShrinkRequest
	GetTagsShrink() *string
}

type UpdateMessageQueueRouteShrinkRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hkhon1po62@c3df23522baa898
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether message canary release is enabled. Valid values:
	//
	// - `true`: Enabled.
	//
	// - `false`: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The filtering side for message canary release.
	//
	// example:
	//
	// Server
	FilterSide *string `json:"FilterSide,omitempty" xml:"FilterSide,omitempty"`
	// The list of baseline environment labels for MSMQ grayscale rules.
	GrayBaseTagsShrink *string `json:"GrayBaseTags,omitempty" xml:"GrayBaseTags,omitempty"`
	// The MSE namespace name.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ignored tags for the untagged environments of the application.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateMessageQueueRouteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageQueueRouteShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetFilterSide() *string {
	return s.FilterSide
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetGrayBaseTagsShrink() *string {
	return s.GrayBaseTagsShrink
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateMessageQueueRouteShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetAcceptLanguage(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetAppId(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetAppName(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetEnable(v bool) *UpdateMessageQueueRouteShrinkRequest {
	s.Enable = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetFilterSide(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.FilterSide = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetGrayBaseTagsShrink(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.GrayBaseTagsShrink = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetNamespace(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetRegion(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.Region = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) SetTagsShrink(v string) *UpdateMessageQueueRouteShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateMessageQueueRouteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
