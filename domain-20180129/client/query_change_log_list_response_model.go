// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChangeLogListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryChangeLogListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryChangeLogListResponse
	GetStatusCode() *int32
	SetBody(v *QueryChangeLogListResponseBody) *QueryChangeLogListResponse
	GetBody() *QueryChangeLogListResponseBody
}

type QueryChangeLogListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChangeLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChangeLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryChangeLogListResponse) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryChangeLogListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryChangeLogListResponse) GetBody() *QueryChangeLogListResponseBody {
	return s.Body
}

func (s *QueryChangeLogListResponse) SetHeaders(v map[string]*string) *QueryChangeLogListResponse {
	s.Headers = v
	return s
}

func (s *QueryChangeLogListResponse) SetStatusCode(v int32) *QueryChangeLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChangeLogListResponse) SetBody(v *QueryChangeLogListResponseBody) *QueryChangeLogListResponse {
	s.Body = v
	return s
}

func (s *QueryChangeLogListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
