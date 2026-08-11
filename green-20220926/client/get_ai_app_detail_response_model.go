// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppDetailResponseBody) *GetAiAppDetailResponse
	GetBody() *GetAiAppDetailResponseBody
}

type GetAiAppDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppDetailResponse) GetBody() *GetAiAppDetailResponseBody {
	return s.Body
}

func (s *GetAiAppDetailResponse) SetHeaders(v map[string]*string) *GetAiAppDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppDetailResponse) SetStatusCode(v int32) *GetAiAppDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppDetailResponse) SetBody(v *GetAiAppDetailResponseBody) *GetAiAppDetailResponse {
	s.Body = v
	return s
}

func (s *GetAiAppDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
