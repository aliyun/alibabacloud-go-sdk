// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKvDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKvDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetKvDetailResponseBody) *GetKvDetailResponse
	GetBody() *GetKvDetailResponseBody
}

type GetKvDetailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKvDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKvDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKvDetailResponse) GoString() string {
	return s.String()
}

func (s *GetKvDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKvDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKvDetailResponse) GetBody() *GetKvDetailResponseBody {
	return s.Body
}

func (s *GetKvDetailResponse) SetHeaders(v map[string]*string) *GetKvDetailResponse {
	s.Headers = v
	return s
}

func (s *GetKvDetailResponse) SetStatusCode(v int32) *GetKvDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKvDetailResponse) SetBody(v *GetKvDetailResponseBody) *GetKvDetailResponse {
	s.Body = v
	return s
}

func (s *GetKvDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
