// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeVisibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeVisibilityResponse
	GetStatusCode() *int32
	SetBody(v *ChangeVisibilityResponseBody) *ChangeVisibilityResponse
	GetBody() *ChangeVisibilityResponseBody
}

type ChangeVisibilityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeVisibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityResponse) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeVisibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeVisibilityResponse) GetBody() *ChangeVisibilityResponseBody {
	return s.Body
}

func (s *ChangeVisibilityResponse) SetHeaders(v map[string]*string) *ChangeVisibilityResponse {
	s.Headers = v
	return s
}

func (s *ChangeVisibilityResponse) SetStatusCode(v int32) *ChangeVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeVisibilityResponse) SetBody(v *ChangeVisibilityResponseBody) *ChangeVisibilityResponse {
	s.Body = v
	return s
}

func (s *ChangeVisibilityResponse) Validate() error {
	return dara.Validate(s)
}
