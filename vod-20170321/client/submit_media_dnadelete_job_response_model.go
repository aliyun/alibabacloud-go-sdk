// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaDNADeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaDNADeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaDNADeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaDNADeleteJobResponseBody) *SubmitMediaDNADeleteJobResponse
	GetBody() *SubmitMediaDNADeleteJobResponseBody
}

type SubmitMediaDNADeleteJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaDNADeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaDNADeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaDNADeleteJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaDNADeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaDNADeleteJobResponse) GetBody() *SubmitMediaDNADeleteJobResponseBody {
	return s.Body
}

func (s *SubmitMediaDNADeleteJobResponse) SetHeaders(v map[string]*string) *SubmitMediaDNADeleteJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaDNADeleteJobResponse) SetStatusCode(v int32) *SubmitMediaDNADeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaDNADeleteJobResponse) SetBody(v *SubmitMediaDNADeleteJobResponseBody) *SubmitMediaDNADeleteJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaDNADeleteJobResponse) Validate() error {
	return dara.Validate(s)
}
