// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppMessageQueueRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetAppMessageQueueRouteRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetAppMessageQueueRouteRequest
	GetAppId() *string
	SetAppName(v string) *GetAppMessageQueueRouteRequest
	GetAppName() *string
	SetNamespace(v string) *GetAppMessageQueueRouteRequest
	GetNamespace() *string
	SetRegion(v string) *GetAppMessageQueueRouteRequest
	GetRegion() *string
}

type GetAppMessageQueueRouteRequest struct {
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
	// The name of the application.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the Microservices Engine (MSE) namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region where the instance resides. Examples:
	//
	// 	- `cn-hangzhou`: China (Hangzhou)
	//
	// 	- `cn-beijing`: China (Beijing)
	//
	// 	- `cn-shanghai`: China (Shanghai)
	//
	// 	- `cn-zhangjiakou`: China (Zhangjiakou)
	//
	// 	- `cn-shenzhen`: China (Shenzhen)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetAppMessageQueueRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppMessageQueueRouteRequest) GoString() string {
	return s.String()
}

func (s *GetAppMessageQueueRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetAppMessageQueueRouteRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAppMessageQueueRouteRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAppMessageQueueRouteRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetAppMessageQueueRouteRequest) GetRegion() *string {
	return s.Region
}

func (s *GetAppMessageQueueRouteRequest) SetAcceptLanguage(v string) *GetAppMessageQueueRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetAppMessageQueueRouteRequest) SetAppId(v string) *GetAppMessageQueueRouteRequest {
	s.AppId = &v
	return s
}

func (s *GetAppMessageQueueRouteRequest) SetAppName(v string) *GetAppMessageQueueRouteRequest {
	s.AppName = &v
	return s
}

func (s *GetAppMessageQueueRouteRequest) SetNamespace(v string) *GetAppMessageQueueRouteRequest {
	s.Namespace = &v
	return s
}

func (s *GetAppMessageQueueRouteRequest) SetRegion(v string) *GetAppMessageQueueRouteRequest {
	s.Region = &v
	return s
}

func (s *GetAppMessageQueueRouteRequest) Validate() error {
	return dara.Validate(s)
}
