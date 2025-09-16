// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostQueryResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPostQueryResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPostQueryResultResponse
	GetStatusCode() *int32
	SetBody(v *ListPostQueryResultResponseBody) *ListPostQueryResultResponse
	GetBody() *ListPostQueryResultResponseBody
}

type ListPostQueryResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPostQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPostQueryResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPostQueryResultResponse) GoString() string {
	return s.String()
}

func (s *ListPostQueryResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPostQueryResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPostQueryResultResponse) GetBody() *ListPostQueryResultResponseBody {
	return s.Body
}

func (s *ListPostQueryResultResponse) SetHeaders(v map[string]*string) *ListPostQueryResultResponse {
	s.Headers = v
	return s
}

func (s *ListPostQueryResultResponse) SetStatusCode(v int32) *ListPostQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPostQueryResultResponse) SetBody(v *ListPostQueryResultResponseBody) *ListPostQueryResultResponse {
	s.Body = v
	return s
}

func (s *ListPostQueryResultResponse) Validate() error {
	return dara.Validate(s)
}
