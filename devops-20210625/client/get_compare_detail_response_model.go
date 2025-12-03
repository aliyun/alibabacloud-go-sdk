// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompareDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompareDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompareDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCompareDetailResponseBody) *GetCompareDetailResponse
	GetBody() *GetCompareDetailResponseBody
}

type GetCompareDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompareDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompareDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCompareDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompareDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompareDetailResponse) GetBody() *GetCompareDetailResponseBody {
	return s.Body
}

func (s *GetCompareDetailResponse) SetHeaders(v map[string]*string) *GetCompareDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCompareDetailResponse) SetStatusCode(v int32) *GetCompareDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompareDetailResponse) SetBody(v *GetCompareDetailResponseBody) *GetCompareDetailResponse {
	s.Body = v
	return s
}

func (s *GetCompareDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
