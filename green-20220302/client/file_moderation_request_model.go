// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *FileModerationRequest
	GetService() *string
	SetServiceParameters(v string) *FileModerationRequest
	GetServiceParameters() *string
}

type FileModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// document_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"url":"https://detect-obj.oss-cn-hangzhou.aliyuncs.com/sample/xxxx.pdf"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s FileModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s FileModerationRequest) GoString() string {
	return s.String()
}

func (s *FileModerationRequest) GetService() *string {
	return s.Service
}

func (s *FileModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *FileModerationRequest) SetService(v string) *FileModerationRequest {
	s.Service = &v
	return s
}

func (s *FileModerationRequest) SetServiceParameters(v string) *FileModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *FileModerationRequest) Validate() error {
	return dara.Validate(s)
}
