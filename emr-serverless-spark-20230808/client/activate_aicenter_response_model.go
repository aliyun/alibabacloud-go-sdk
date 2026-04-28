// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateAICenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateAICenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateAICenterResponse
	GetStatusCode() *int32
	SetBody(v *ActivateAICenterResponseBody) *ActivateAICenterResponse
	GetBody() *ActivateAICenterResponseBody
}

type ActivateAICenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateAICenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateAICenterResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateAICenterResponse) GoString() string {
	return s.String()
}

func (s *ActivateAICenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateAICenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateAICenterResponse) GetBody() *ActivateAICenterResponseBody {
	return s.Body
}

func (s *ActivateAICenterResponse) SetHeaders(v map[string]*string) *ActivateAICenterResponse {
	s.Headers = v
	return s
}

func (s *ActivateAICenterResponse) SetStatusCode(v int32) *ActivateAICenterResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateAICenterResponse) SetBody(v *ActivateAICenterResponseBody) *ActivateAICenterResponse {
	s.Body = v
	return s
}

func (s *ActivateAICenterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
