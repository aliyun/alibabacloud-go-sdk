// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYikeUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYikeUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateYikeUserResponseBody) *CreateYikeUserResponse
	GetBody() *CreateYikeUserResponseBody
}

type CreateYikeUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYikeUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYikeUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeUserResponse) GoString() string {
	return s.String()
}

func (s *CreateYikeUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYikeUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYikeUserResponse) GetBody() *CreateYikeUserResponseBody {
	return s.Body
}

func (s *CreateYikeUserResponse) SetHeaders(v map[string]*string) *CreateYikeUserResponse {
	s.Headers = v
	return s
}

func (s *CreateYikeUserResponse) SetStatusCode(v int32) *CreateYikeUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYikeUserResponse) SetBody(v *CreateYikeUserResponseBody) *CreateYikeUserResponse {
	s.Body = v
	return s
}

func (s *CreateYikeUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
