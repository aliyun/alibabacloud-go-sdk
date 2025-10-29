// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveAudioAuditConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveAudioAuditConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveAudioAuditConfigResponseBody) *DescribeLiveAudioAuditConfigResponse
	GetBody() *DescribeLiveAudioAuditConfigResponseBody
}

type DescribeLiveAudioAuditConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveAudioAuditConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveAudioAuditConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveAudioAuditConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveAudioAuditConfigResponse) GetBody() *DescribeLiveAudioAuditConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveAudioAuditConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveAudioAuditConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponse) SetStatusCode(v int32) *DescribeLiveAudioAuditConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponse) SetBody(v *DescribeLiveAudioAuditConfigResponseBody) *DescribeLiveAudioAuditConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
