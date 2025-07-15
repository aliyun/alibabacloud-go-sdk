// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveAudioAuditNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveAudioAuditNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveAudioAuditNotifyConfigResponseBody) *DescribeLiveAudioAuditNotifyConfigResponse
	GetBody() *DescribeLiveAudioAuditNotifyConfigResponseBody
}

type DescribeLiveAudioAuditNotifyConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveAudioAuditNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) GetBody() *DescribeLiveAudioAuditNotifyConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveAudioAuditNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) SetStatusCode(v int32) *DescribeLiveAudioAuditNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) SetBody(v *DescribeLiveAudioAuditNotifyConfigResponseBody) *DescribeLiveAudioAuditNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
