// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduledPreloadExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduledPreloadExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduledPreloadExecutionsResponseBody) *CreateScheduledPreloadExecutionsResponse
	GetBody() *CreateScheduledPreloadExecutionsResponseBody
}

type CreateScheduledPreloadExecutionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPreloadExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduledPreloadExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduledPreloadExecutionsResponse) GetBody() *CreateScheduledPreloadExecutionsResponseBody {
	return s.Body
}

func (s *CreateScheduledPreloadExecutionsResponse) SetHeaders(v map[string]*string) *CreateScheduledPreloadExecutionsResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponse) SetStatusCode(v int32) *CreateScheduledPreloadExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponse) SetBody(v *CreateScheduledPreloadExecutionsResponseBody) *CreateScheduledPreloadExecutionsResponse {
	s.Body = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponse) Validate() error {
	return dara.Validate(s)
}
