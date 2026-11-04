// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserIdByCustomUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserIdByCustomUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserIdByCustomUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserIdByCustomUserIdResponseBody) *QueryFaceUserIdByCustomUserIdResponse
	GetBody() *QueryFaceUserIdByCustomUserIdResponseBody
}

type QueryFaceUserIdByCustomUserIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserIdByCustomUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserIdByCustomUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserIdByCustomUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserIdByCustomUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserIdByCustomUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserIdByCustomUserIdResponse) GetBody() *QueryFaceUserIdByCustomUserIdResponseBody {
	return s.Body
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetHeaders(v map[string]*string) *QueryFaceUserIdByCustomUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetStatusCode(v int32) *QueryFaceUserIdByCustomUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponse) SetBody(v *QueryFaceUserIdByCustomUserIdResponseBody) *QueryFaceUserIdByCustomUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserIdByCustomUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
