// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageQueueRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateMessageQueueRouteRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *UpdateMessageQueueRouteRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMessageQueueRouteRequest
	GetAppName() *string
	SetEnable(v bool) *UpdateMessageQueueRouteRequest
	GetEnable() *bool
	SetFilterSide(v string) *UpdateMessageQueueRouteRequest
	GetFilterSide() *string
	SetGrayBaseTags(v []*string) *UpdateMessageQueueRouteRequest
	GetGrayBaseTags() []*string
	SetNamespace(v string) *UpdateMessageQueueRouteRequest
	GetNamespace() *string
	SetRegion(v string) *UpdateMessageQueueRouteRequest
	GetRegion() *string
	SetTags(v []*string) *UpdateMessageQueueRouteRequest
	GetTags() []*string
}

type UpdateMessageQueueRouteRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// hkhon1po62@c3df23522baa898
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether the canary release for messaging feature is enabled for the application. Valid values:
	//
	// 	- `true`: enabled
	//
	// 	- `false`: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The side for message filtering when the canary release for messaging feature is enabled.
	//
	// example:
	//
	// Server
	FilterSide   *string   `json:"FilterSide,omitempty" xml:"FilterSide,omitempty"`
	GrayBaseTags []*string `json:"GrayBaseTags,omitempty" xml:"GrayBaseTags,omitempty" type:"Repeated"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The tag that is negligible for the untagged environment of the application.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateMessageQueueRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageQueueRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageQueueRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateMessageQueueRouteRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageQueueRouteRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMessageQueueRouteRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateMessageQueueRouteRequest) GetFilterSide() *string {
	return s.FilterSide
}

func (s *UpdateMessageQueueRouteRequest) GetGrayBaseTags() []*string {
	return s.GrayBaseTags
}

func (s *UpdateMessageQueueRouteRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateMessageQueueRouteRequest) GetRegion() *string {
	return s.Region
}

func (s *UpdateMessageQueueRouteRequest) GetTags() []*string {
	return s.Tags
}

func (s *UpdateMessageQueueRouteRequest) SetAcceptLanguage(v string) *UpdateMessageQueueRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetAppId(v string) *UpdateMessageQueueRouteRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetAppName(v string) *UpdateMessageQueueRouteRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetEnable(v bool) *UpdateMessageQueueRouteRequest {
	s.Enable = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetFilterSide(v string) *UpdateMessageQueueRouteRequest {
	s.FilterSide = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetGrayBaseTags(v []*string) *UpdateMessageQueueRouteRequest {
	s.GrayBaseTags = v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetNamespace(v string) *UpdateMessageQueueRouteRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetRegion(v string) *UpdateMessageQueueRouteRequest {
	s.Region = &v
	return s
}

func (s *UpdateMessageQueueRouteRequest) SetTags(v []*string) *UpdateMessageQueueRouteRequest {
	s.Tags = v
	return s
}

func (s *UpdateMessageQueueRouteRequest) Validate() error {
	return dara.Validate(s)
}
