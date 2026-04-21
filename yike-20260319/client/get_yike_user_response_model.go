// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeUserResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeUserResponseBody) *GetYikeUserResponse
	GetBody() *GetYikeUserResponseBody
}

type GetYikeUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeUserResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserResponse) GoString() string {
	return s.String()
}

func (s *GetYikeUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeUserResponse) GetBody() *GetYikeUserResponseBody {
	return s.Body
}

func (s *GetYikeUserResponse) SetHeaders(v map[string]*string) *GetYikeUserResponse {
	s.Headers = v
	return s
}

func (s *GetYikeUserResponse) SetStatusCode(v int32) *GetYikeUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeUserResponse) SetBody(v *GetYikeUserResponseBody) *GetYikeUserResponse {
	s.Body = v
	return s
}

func (s *GetYikeUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
