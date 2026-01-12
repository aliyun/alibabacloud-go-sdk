// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProjectDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProjectDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryProjectDetailResponseBody) *QueryProjectDetailResponse
	GetBody() *QueryProjectDetailResponseBody
}

type QueryProjectDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProjectDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProjectDetailResponse) GetBody() *QueryProjectDetailResponseBody {
	return s.Body
}

func (s *QueryProjectDetailResponse) SetHeaders(v map[string]*string) *QueryProjectDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectDetailResponse) SetStatusCode(v int32) *QueryProjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectDetailResponse) SetBody(v *QueryProjectDetailResponseBody) *QueryProjectDetailResponse {
	s.Body = v
	return s
}

func (s *QueryProjectDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
