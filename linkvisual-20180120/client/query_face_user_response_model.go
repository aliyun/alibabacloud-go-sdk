// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserResponseBody) *QueryFaceUserResponse
	GetBody() *QueryFaceUserResponseBody
}

type QueryFaceUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserResponse) GetBody() *QueryFaceUserResponseBody {
	return s.Body
}

func (s *QueryFaceUserResponse) SetHeaders(v map[string]*string) *QueryFaceUserResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserResponse) SetStatusCode(v int32) *QueryFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserResponse) SetBody(v *QueryFaceUserResponseBody) *QueryFaceUserResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
