// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *ShareLink) *UpdateShareLinkResponse
	GetBody() *ShareLink
}

type UpdateShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateShareLinkResponse) GetBody() *ShareLink {
	return s.Body
}

func (s *UpdateShareLinkResponse) SetHeaders(v map[string]*string) *UpdateShareLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateShareLinkResponse) SetStatusCode(v int32) *UpdateShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShareLinkResponse) SetBody(v *ShareLink) *UpdateShareLinkResponse {
	s.Body = v
	return s
}

func (s *UpdateShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
