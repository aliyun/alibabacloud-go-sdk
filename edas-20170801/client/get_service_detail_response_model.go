// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceDetailResponseBody) *GetServiceDetailResponse
	GetBody() *GetServiceDetailResponseBody
}

type GetServiceDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceDetailResponse) GetBody() *GetServiceDetailResponseBody {
	return s.Body
}

func (s *GetServiceDetailResponse) SetHeaders(v map[string]*string) *GetServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetServiceDetailResponse) SetStatusCode(v int32) *GetServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceDetailResponse) SetBody(v *GetServiceDetailResponseBody) *GetServiceDetailResponse {
	s.Body = v
	return s
}

func (s *GetServiceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
