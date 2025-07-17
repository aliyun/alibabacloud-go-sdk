// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoInsightsActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DoInsightsActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DoInsightsActionResponse
	GetStatusCode() *int32
	SetBody(v *DoInsightsActionResponseBody) *DoInsightsActionResponse
	GetBody() *DoInsightsActionResponseBody
}

type DoInsightsActionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DoInsightsActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DoInsightsActionResponse) String() string {
	return dara.Prettify(s)
}

func (s DoInsightsActionResponse) GoString() string {
	return s.String()
}

func (s *DoInsightsActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DoInsightsActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DoInsightsActionResponse) GetBody() *DoInsightsActionResponseBody {
	return s.Body
}

func (s *DoInsightsActionResponse) SetHeaders(v map[string]*string) *DoInsightsActionResponse {
	s.Headers = v
	return s
}

func (s *DoInsightsActionResponse) SetStatusCode(v int32) *DoInsightsActionResponse {
	s.StatusCode = &v
	return s
}

func (s *DoInsightsActionResponse) SetBody(v *DoInsightsActionResponseBody) *DoInsightsActionResponse {
	s.Body = v
	return s
}

func (s *DoInsightsActionResponse) Validate() error {
	return dara.Validate(s)
}
