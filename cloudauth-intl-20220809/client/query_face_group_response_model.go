// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceGroupResponseBody) *QueryFaceGroupResponse
	GetBody() *QueryFaceGroupResponseBody
}

type QueryFaceGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceGroupResponse) GetBody() *QueryFaceGroupResponseBody {
	return s.Body
}

func (s *QueryFaceGroupResponse) SetHeaders(v map[string]*string) *QueryFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceGroupResponse) SetStatusCode(v int32) *QueryFaceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceGroupResponse) SetBody(v *QueryFaceGroupResponseBody) *QueryFaceGroupResponse {
	s.Body = v
	return s
}

func (s *QueryFaceGroupResponse) Validate() error {
	return dara.Validate(s)
}
