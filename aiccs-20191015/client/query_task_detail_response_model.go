// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskDetailResponseBody) *QueryTaskDetailResponse
	GetBody() *QueryTaskDetailResponseBody
}

type QueryTaskDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskDetailResponse) GetBody() *QueryTaskDetailResponseBody {
	return s.Body
}

func (s *QueryTaskDetailResponse) SetHeaders(v map[string]*string) *QueryTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailResponse) SetStatusCode(v int32) *QueryTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskDetailResponse) SetBody(v *QueryTaskDetailResponseBody) *QueryTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
