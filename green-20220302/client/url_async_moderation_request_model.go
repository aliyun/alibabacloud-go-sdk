// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUrlAsyncModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *UrlAsyncModerationRequest
	GetService() *string
	SetServiceParameters(v string) *UrlAsyncModerationRequest
	GetServiceParameters() *string
}

type UrlAsyncModerationRequest struct {
	// Service name: URL threat detection
	//
	// example:
	//
	// url_detection_pro
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set for the content moderation object. This parameter is a JSON string. For more information, see the description of ServiceParameters.
	//
	// example:
	//
	// {
	//
	//         "url": "https://help.aliyun.com/",
	//
	//         "dataId": "url123******"
	//
	// }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s UrlAsyncModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s UrlAsyncModerationRequest) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationRequest) GetService() *string {
	return s.Service
}

func (s *UrlAsyncModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *UrlAsyncModerationRequest) SetService(v string) *UrlAsyncModerationRequest {
	s.Service = &v
	return s
}

func (s *UrlAsyncModerationRequest) SetServiceParameters(v string) *UrlAsyncModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *UrlAsyncModerationRequest) Validate() error {
	return dara.Validate(s)
}
