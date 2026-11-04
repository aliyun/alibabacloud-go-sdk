// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserIdsByGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceAllUserIdsByGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceAllUserIdsByGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceAllUserIdsByGroupIdResponseBody) *QueryFaceAllUserIdsByGroupIdResponse
	GetBody() *QueryFaceAllUserIdsByGroupIdResponseBody
}

type QueryFaceAllUserIdsByGroupIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserIdsByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) GetBody() *QueryFaceAllUserIdsByGroupIdResponseBody {
	return s.Body
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserIdsByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetStatusCode(v int32) *QueryFaceAllUserIdsByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) SetBody(v *QueryFaceAllUserIdsByGroupIdResponseBody) *QueryFaceAllUserIdsByGroupIdResponse {
	s.Body = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
