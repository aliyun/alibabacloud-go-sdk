// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VideoModerationRequest
	GetService() *string
	SetServiceParameters(v string) *VideoModerationRequest
	GetServiceParameters() *string
}

type VideoModerationRequest struct {
	// The service code for video moderation.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters that are required for the moderation service. The value must be a JSON string.
	//
	// - url: Required. The URL of the object to be moderated. Make sure that the URL can be accessed over the Internet.
	//
	// - dataId: Optional. The data ID of the object to be moderated.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/2505810.html).
	//
	// example:
	//
	// {"url": "https://talesofai.oss-cn-shanghai.aliyuncs.com/xxx.mp4", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationRequest) GetService() *string {
	return s.Service
}

func (s *VideoModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VideoModerationRequest) SetService(v string) *VideoModerationRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationRequest) SetServiceParameters(v string) *VideoModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VideoModerationRequest) Validate() error {
	return dara.Validate(s)
}
