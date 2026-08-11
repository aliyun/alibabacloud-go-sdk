// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppNodeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppNodeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppNodeDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppNodeDetailResponseBody) *GetAiAppNodeDetailResponse
	GetBody() *GetAiAppNodeDetailResponseBody
}

type GetAiAppNodeDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppNodeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppNodeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppNodeDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppNodeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppNodeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppNodeDetailResponse) GetBody() *GetAiAppNodeDetailResponseBody {
	return s.Body
}

func (s *GetAiAppNodeDetailResponse) SetHeaders(v map[string]*string) *GetAiAppNodeDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppNodeDetailResponse) SetStatusCode(v int32) *GetAiAppNodeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppNodeDetailResponse) SetBody(v *GetAiAppNodeDetailResponseBody) *GetAiAppNodeDetailResponse {
	s.Body = v
	return s
}

func (s *GetAiAppNodeDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
