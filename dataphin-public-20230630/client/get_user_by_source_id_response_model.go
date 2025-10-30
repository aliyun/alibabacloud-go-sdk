// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBySourceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserBySourceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserBySourceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserBySourceIdResponseBody) *GetUserBySourceIdResponse
	GetBody() *GetUserBySourceIdResponseBody
}

type GetUserBySourceIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBySourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBySourceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserBySourceIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserBySourceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserBySourceIdResponse) GetBody() *GetUserBySourceIdResponseBody {
	return s.Body
}

func (s *GetUserBySourceIdResponse) SetHeaders(v map[string]*string) *GetUserBySourceIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserBySourceIdResponse) SetStatusCode(v int32) *GetUserBySourceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBySourceIdResponse) SetBody(v *GetUserBySourceIdResponseBody) *GetUserBySourceIdResponse {
	s.Body = v
	return s
}

func (s *GetUserBySourceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
