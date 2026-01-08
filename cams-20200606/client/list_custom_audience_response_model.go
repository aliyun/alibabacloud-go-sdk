// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAudienceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAudienceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAudienceResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAudienceResponseBody) *ListCustomAudienceResponse
	GetBody() *ListCustomAudienceResponseBody
}

type ListCustomAudienceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAudienceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAudienceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAudienceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAudienceResponse) GetBody() *ListCustomAudienceResponseBody {
	return s.Body
}

func (s *ListCustomAudienceResponse) SetHeaders(v map[string]*string) *ListCustomAudienceResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAudienceResponse) SetStatusCode(v int32) *ListCustomAudienceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAudienceResponse) SetBody(v *ListCustomAudienceResponseBody) *ListCustomAudienceResponse {
	s.Body = v
	return s
}

func (s *ListCustomAudienceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
