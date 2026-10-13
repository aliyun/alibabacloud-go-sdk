// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePortalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePortalResponse
	GetStatusCode() *int32
	SetBody(v *CreatePortalResponseBody) *CreatePortalResponse
	GetBody() *CreatePortalResponseBody
}

type CreatePortalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePortalResponse) GoString() string {
	return s.String()
}

func (s *CreatePortalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePortalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePortalResponse) GetBody() *CreatePortalResponseBody {
	return s.Body
}

func (s *CreatePortalResponse) SetHeaders(v map[string]*string) *CreatePortalResponse {
	s.Headers = v
	return s
}

func (s *CreatePortalResponse) SetStatusCode(v int32) *CreatePortalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortalResponse) SetBody(v *CreatePortalResponseBody) *CreatePortalResponse {
	s.Body = v
	return s
}

func (s *CreatePortalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
