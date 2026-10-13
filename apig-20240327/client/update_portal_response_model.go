// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePortalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePortalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePortalResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePortalResponseBody) *UpdatePortalResponse
	GetBody() *UpdatePortalResponseBody
}

type UpdatePortalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePortalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePortalResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortalResponse) GoString() string {
	return s.String()
}

func (s *UpdatePortalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePortalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePortalResponse) GetBody() *UpdatePortalResponseBody {
	return s.Body
}

func (s *UpdatePortalResponse) SetHeaders(v map[string]*string) *UpdatePortalResponse {
	s.Headers = v
	return s
}

func (s *UpdatePortalResponse) SetStatusCode(v int32) *UpdatePortalResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePortalResponse) SetBody(v *UpdatePortalResponseBody) *UpdatePortalResponse {
	s.Body = v
	return s
}

func (s *UpdatePortalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
