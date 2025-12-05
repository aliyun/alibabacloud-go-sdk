// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserVpcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserVpcsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserVpcsResponseBody) *GetUserVpcsResponse
	GetBody() *GetUserVpcsResponseBody
}

type GetUserVpcsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserVpcsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcsResponse) GoString() string {
	return s.String()
}

func (s *GetUserVpcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserVpcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserVpcsResponse) GetBody() *GetUserVpcsResponseBody {
	return s.Body
}

func (s *GetUserVpcsResponse) SetHeaders(v map[string]*string) *GetUserVpcsResponse {
	s.Headers = v
	return s
}

func (s *GetUserVpcsResponse) SetStatusCode(v int32) *GetUserVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserVpcsResponse) SetBody(v *GetUserVpcsResponseBody) *GetUserVpcsResponse {
	s.Body = v
	return s
}

func (s *GetUserVpcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
