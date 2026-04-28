// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *ShareLink) *CreateShareLinkResponse
	GetBody() *ShareLink
}

type CreateShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateShareLinkResponse) GetBody() *ShareLink {
	return s.Body
}

func (s *CreateShareLinkResponse) SetHeaders(v map[string]*string) *CreateShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateShareLinkResponse) SetStatusCode(v int32) *CreateShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShareLinkResponse) SetBody(v *ShareLink) *CreateShareLinkResponse {
	s.Body = v
	return s
}

func (s *CreateShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
