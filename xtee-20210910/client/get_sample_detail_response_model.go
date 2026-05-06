// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSampleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSampleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSampleDetailResponseBody) *GetSampleDetailResponse
	GetBody() *GetSampleDetailResponseBody
}

type GetSampleDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSampleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSampleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSampleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSampleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSampleDetailResponse) GetBody() *GetSampleDetailResponseBody {
	return s.Body
}

func (s *GetSampleDetailResponse) SetHeaders(v map[string]*string) *GetSampleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSampleDetailResponse) SetStatusCode(v int32) *GetSampleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSampleDetailResponse) SetBody(v *GetSampleDetailResponseBody) *GetSampleDetailResponse {
	s.Body = v
	return s
}

func (s *GetSampleDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
