// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCdsFileShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCdsFileShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateCdsFileShareLinkResponseBody) *CreateCdsFileShareLinkResponse
	GetBody() *CreateCdsFileShareLinkResponseBody
}

type CreateCdsFileShareLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdsFileShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdsFileShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateCdsFileShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCdsFileShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCdsFileShareLinkResponse) GetBody() *CreateCdsFileShareLinkResponseBody {
	return s.Body
}

func (s *CreateCdsFileShareLinkResponse) SetHeaders(v map[string]*string) *CreateCdsFileShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateCdsFileShareLinkResponse) SetStatusCode(v int32) *CreateCdsFileShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdsFileShareLinkResponse) SetBody(v *CreateCdsFileShareLinkResponseBody) *CreateCdsFileShareLinkResponse {
	s.Body = v
	return s
}

func (s *CreateCdsFileShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
