// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCaseDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCaseDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCaseDetailResponseBody) *GetCaseDetailResponse
	GetBody() *GetCaseDetailResponseBody
}

type GetCaseDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaseDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCaseDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCaseDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCaseDetailResponse) GetBody() *GetCaseDetailResponseBody {
	return s.Body
}

func (s *GetCaseDetailResponse) SetHeaders(v map[string]*string) *GetCaseDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCaseDetailResponse) SetStatusCode(v int32) *GetCaseDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaseDetailResponse) SetBody(v *GetCaseDetailResponseBody) *GetCaseDetailResponse {
	s.Body = v
	return s
}

func (s *GetCaseDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
