// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByOpenDingtalkIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByOpenDingtalkIdResponseBody) *GetUserIdByOpenDingtalkIdResponse
	GetBody() *GetUserIdByOpenDingtalkIdResponseBody
}

type GetUserIdByOpenDingtalkIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByOpenDingtalkIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByOpenDingtalkIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByOpenDingtalkIdResponse) GetBody() *GetUserIdByOpenDingtalkIdResponseBody {
	return s.Body
}

func (s *GetUserIdByOpenDingtalkIdResponse) SetHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponse) SetStatusCode(v int32) *GetUserIdByOpenDingtalkIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponse) SetBody(v *GetUserIdByOpenDingtalkIdResponseBody) *GetUserIdByOpenDingtalkIdResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
