// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceCustomUserIdByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceCustomUserIdByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceCustomUserIdByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceCustomUserIdByUserIdResponseBody) *QueryFaceCustomUserIdByUserIdResponse
	GetBody() *QueryFaceCustomUserIdByUserIdResponseBody
}

type QueryFaceCustomUserIdByUserIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceCustomUserIdByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceCustomUserIdByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceCustomUserIdByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceCustomUserIdByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceCustomUserIdByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceCustomUserIdByUserIdResponse) GetBody() *QueryFaceCustomUserIdByUserIdResponseBody {
	return s.Body
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetHeaders(v map[string]*string) *QueryFaceCustomUserIdByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetStatusCode(v int32) *QueryFaceCustomUserIdByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponse) SetBody(v *QueryFaceCustomUserIdByUserIdResponseBody) *QueryFaceCustomUserIdByUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryFaceCustomUserIdByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
