// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryDetailResponseBody) *GetMemoryDetailResponse
	GetBody() *GetMemoryDetailResponseBody
}

type GetMemoryDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryDetailResponse) GetBody() *GetMemoryDetailResponseBody {
	return s.Body
}

func (s *GetMemoryDetailResponse) SetHeaders(v map[string]*string) *GetMemoryDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryDetailResponse) SetStatusCode(v int32) *GetMemoryDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryDetailResponse) SetBody(v *GetMemoryDetailResponseBody) *GetMemoryDetailResponse {
	s.Body = v
	return s
}

func (s *GetMemoryDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
