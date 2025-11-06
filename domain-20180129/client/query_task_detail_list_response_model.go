// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskDetailListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskDetailListResponseBody) *QueryTaskDetailListResponse
	GetBody() *QueryTaskDetailListResponseBody
}

type QueryTaskDetailListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskDetailListResponse) GetBody() *QueryTaskDetailListResponseBody {
	return s.Body
}

func (s *QueryTaskDetailListResponse) SetHeaders(v map[string]*string) *QueryTaskDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailListResponse) SetStatusCode(v int32) *QueryTaskDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskDetailListResponse) SetBody(v *QueryTaskDetailListResponseBody) *QueryTaskDetailListResponse {
	s.Body = v
	return s
}

func (s *QueryTaskDetailListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
