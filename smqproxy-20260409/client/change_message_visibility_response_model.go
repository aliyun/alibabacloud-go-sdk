// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMessageVisibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeMessageVisibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeMessageVisibilityResponse
	GetStatusCode() *int32
	SetBody(v *ChangeMessageVisibilityResponseBody) *ChangeMessageVisibilityResponse
	GetBody() *ChangeMessageVisibilityResponseBody
}

type ChangeMessageVisibilityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMessageVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMessageVisibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeMessageVisibilityResponse) GoString() string {
	return s.String()
}

func (s *ChangeMessageVisibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeMessageVisibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeMessageVisibilityResponse) GetBody() *ChangeMessageVisibilityResponseBody {
	return s.Body
}

func (s *ChangeMessageVisibilityResponse) SetHeaders(v map[string]*string) *ChangeMessageVisibilityResponse {
	s.Headers = v
	return s
}

func (s *ChangeMessageVisibilityResponse) SetStatusCode(v int32) *ChangeMessageVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMessageVisibilityResponse) SetBody(v *ChangeMessageVisibilityResponseBody) *ChangeMessageVisibilityResponse {
	s.Body = v
	return s
}

func (s *ChangeMessageVisibilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
