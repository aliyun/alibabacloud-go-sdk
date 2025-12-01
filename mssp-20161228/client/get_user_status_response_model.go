// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetUserStatusResponseBody) *GetUserStatusResponse
	GetBody() *GetUserStatusResponseBody
}

type GetUserStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserStatusResponse) GetBody() *GetUserStatusResponseBody {
	return s.Body
}

func (s *GetUserStatusResponse) SetHeaders(v map[string]*string) *GetUserStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserStatusResponse) SetStatusCode(v int32) *GetUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserStatusResponse) SetBody(v *GetUserStatusResponseBody) *GetUserStatusResponse {
	s.Body = v
	return s
}

func (s *GetUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
