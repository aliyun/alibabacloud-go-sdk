// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallNumberDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CallNumberDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CallNumberDetailResponse
	GetStatusCode() *int32
	SetBody(v *CallNumberDetailResponseBody) *CallNumberDetailResponse
	GetBody() *CallNumberDetailResponseBody
}

type CallNumberDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallNumberDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallNumberDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s CallNumberDetailResponse) GoString() string {
	return s.String()
}

func (s *CallNumberDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CallNumberDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CallNumberDetailResponse) GetBody() *CallNumberDetailResponseBody {
	return s.Body
}

func (s *CallNumberDetailResponse) SetHeaders(v map[string]*string) *CallNumberDetailResponse {
	s.Headers = v
	return s
}

func (s *CallNumberDetailResponse) SetStatusCode(v int32) *CallNumberDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *CallNumberDetailResponse) SetBody(v *CallNumberDetailResponseBody) *CallNumberDetailResponse {
	s.Body = v
	return s
}

func (s *CallNumberDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
