// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProjectListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProjectListResponse
	GetStatusCode() *int32
	SetBody(v *QueryProjectListResponseBody) *QueryProjectListResponse
	GetBody() *QueryProjectListResponseBody
}

type QueryProjectListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProjectListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProjectListResponse) GetBody() *QueryProjectListResponseBody {
	return s.Body
}

func (s *QueryProjectListResponse) SetHeaders(v map[string]*string) *QueryProjectListResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectListResponse) SetStatusCode(v int32) *QueryProjectListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectListResponse) SetBody(v *QueryProjectListResponseBody) *QueryProjectListResponse {
	s.Body = v
	return s
}

func (s *QueryProjectListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
