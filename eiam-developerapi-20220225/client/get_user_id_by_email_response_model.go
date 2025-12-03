// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByEmailResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByEmailResponseBody) *GetUserIdByEmailResponse
	GetBody() *GetUserIdByEmailResponseBody
}

type GetUserIdByEmailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByEmailResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByEmailResponse) GetBody() *GetUserIdByEmailResponseBody {
	return s.Body
}

func (s *GetUserIdByEmailResponse) SetHeaders(v map[string]*string) *GetUserIdByEmailResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByEmailResponse) SetStatusCode(v int32) *GetUserIdByEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByEmailResponse) SetBody(v *GetUserIdByEmailResponseBody) *GetUserIdByEmailResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
