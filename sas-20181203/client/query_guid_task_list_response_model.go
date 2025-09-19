// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGuidTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGuidTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGuidTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryGuidTaskListResponseBody) *QueryGuidTaskListResponse
	GetBody() *QueryGuidTaskListResponseBody
}

type QueryGuidTaskListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGuidTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGuidTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGuidTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryGuidTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGuidTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGuidTaskListResponse) GetBody() *QueryGuidTaskListResponseBody {
	return s.Body
}

func (s *QueryGuidTaskListResponse) SetHeaders(v map[string]*string) *QueryGuidTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryGuidTaskListResponse) SetStatusCode(v int32) *QueryGuidTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGuidTaskListResponse) SetBody(v *QueryGuidTaskListResponseBody) *QueryGuidTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryGuidTaskListResponse) Validate() error {
	return dara.Validate(s)
}
