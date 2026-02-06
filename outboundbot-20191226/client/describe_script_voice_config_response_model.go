// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptVoiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScriptVoiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScriptVoiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScriptVoiceConfigResponseBody) *DescribeScriptVoiceConfigResponse
	GetBody() *DescribeScriptVoiceConfigResponseBody
}

type DescribeScriptVoiceConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScriptVoiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScriptVoiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptVoiceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeScriptVoiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScriptVoiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScriptVoiceConfigResponse) GetBody() *DescribeScriptVoiceConfigResponseBody {
	return s.Body
}

func (s *DescribeScriptVoiceConfigResponse) SetHeaders(v map[string]*string) *DescribeScriptVoiceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeScriptVoiceConfigResponse) SetStatusCode(v int32) *DescribeScriptVoiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScriptVoiceConfigResponse) SetBody(v *DescribeScriptVoiceConfigResponseBody) *DescribeScriptVoiceConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeScriptVoiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
