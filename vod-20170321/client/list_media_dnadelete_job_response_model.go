// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNADeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaDNADeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaDNADeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaDNADeleteJobResponseBody) *ListMediaDNADeleteJobResponse
	GetBody() *ListMediaDNADeleteJobResponseBody
}

type ListMediaDNADeleteJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaDNADeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaDNADeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNADeleteJobResponse) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaDNADeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaDNADeleteJobResponse) GetBody() *ListMediaDNADeleteJobResponseBody {
	return s.Body
}

func (s *ListMediaDNADeleteJobResponse) SetHeaders(v map[string]*string) *ListMediaDNADeleteJobResponse {
	s.Headers = v
	return s
}

func (s *ListMediaDNADeleteJobResponse) SetStatusCode(v int32) *ListMediaDNADeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaDNADeleteJobResponse) SetBody(v *ListMediaDNADeleteJobResponseBody) *ListMediaDNADeleteJobResponse {
	s.Body = v
	return s
}

func (s *ListMediaDNADeleteJobResponse) Validate() error {
	return dara.Validate(s)
}
