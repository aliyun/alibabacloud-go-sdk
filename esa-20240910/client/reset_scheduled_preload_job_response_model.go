// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetScheduledPreloadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetScheduledPreloadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetScheduledPreloadJobResponse
	GetStatusCode() *int32
	SetBody(v *ResetScheduledPreloadJobResponseBody) *ResetScheduledPreloadJobResponse
	GetBody() *ResetScheduledPreloadJobResponseBody
}

type ResetScheduledPreloadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetScheduledPreloadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetScheduledPreloadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetScheduledPreloadJobResponse) GetBody() *ResetScheduledPreloadJobResponseBody {
	return s.Body
}

func (s *ResetScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *ResetScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *ResetScheduledPreloadJobResponse) SetStatusCode(v int32) *ResetScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetScheduledPreloadJobResponse) SetBody(v *ResetScheduledPreloadJobResponseBody) *ResetScheduledPreloadJobResponse {
	s.Body = v
	return s
}

func (s *ResetScheduledPreloadJobResponse) Validate() error {
	return dara.Validate(s)
}
