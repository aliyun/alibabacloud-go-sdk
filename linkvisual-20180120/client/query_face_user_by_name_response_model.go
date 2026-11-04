// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserByNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserByNameResponseBody) *QueryFaceUserByNameResponse
	GetBody() *QueryFaceUserByNameResponseBody
}

type QueryFaceUserByNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserByNameResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserByNameResponse) GetBody() *QueryFaceUserByNameResponseBody {
	return s.Body
}

func (s *QueryFaceUserByNameResponse) SetHeaders(v map[string]*string) *QueryFaceUserByNameResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserByNameResponse) SetStatusCode(v int32) *QueryFaceUserByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserByNameResponse) SetBody(v *QueryFaceUserByNameResponseBody) *QueryFaceUserByNameResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
