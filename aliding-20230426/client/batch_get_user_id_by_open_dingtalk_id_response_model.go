// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetUserIdByOpenDingtalkIdResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetUserIdByOpenDingtalkIdResponseBody) *BatchGetUserIdByOpenDingtalkIdResponse
	GetBody() *BatchGetUserIdByOpenDingtalkIdResponseBody
}

type BatchGetUserIdByOpenDingtalkIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetUserIdByOpenDingtalkIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdResponse) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) GetBody() *BatchGetUserIdByOpenDingtalkIdResponseBody {
	return s.Body
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) SetHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdResponse {
	s.Headers = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) SetStatusCode(v int32) *BatchGetUserIdByOpenDingtalkIdResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) SetBody(v *BatchGetUserIdByOpenDingtalkIdResponseBody) *BatchGetUserIdByOpenDingtalkIdResponse {
	s.Body = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
