// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *DescribeFileModerationResultRequest
	GetService() *string
	SetServiceParameters(v string) *DescribeFileModerationResultRequest
	GetServiceParameters() *string
}

type DescribeFileModerationResultRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// document_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// 	- taskId: required. The URL of the object that you want to moderate. Make sure that the URL can be accessed over the Internet.
	//
	// example:
	//
	// {\\"taskId\\":\\"vi_f_hPgx9PFIQISdlfA888hOFG-1yJq8v\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s DescribeFileModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultRequest) GetService() *string {
	return s.Service
}

func (s *DescribeFileModerationResultRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *DescribeFileModerationResultRequest) SetService(v string) *DescribeFileModerationResultRequest {
	s.Service = &v
	return s
}

func (s *DescribeFileModerationResultRequest) SetServiceParameters(v string) *DescribeFileModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

func (s *DescribeFileModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
