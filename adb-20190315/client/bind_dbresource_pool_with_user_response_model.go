// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourcePoolWithUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindDBResourcePoolWithUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindDBResourcePoolWithUserResponse
	GetStatusCode() *int32
	SetBody(v *BindDBResourcePoolWithUserResponseBody) *BindDBResourcePoolWithUserResponse
	GetBody() *BindDBResourcePoolWithUserResponseBody
}

type BindDBResourcePoolWithUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDBResourcePoolWithUserResponse) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindDBResourcePoolWithUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindDBResourcePoolWithUserResponse) GetBody() *BindDBResourcePoolWithUserResponseBody {
	return s.Body
}

func (s *BindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) SetStatusCode(v int32) *BindDBResourcePoolWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) SetBody(v *BindDBResourcePoolWithUserResponseBody) *BindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
