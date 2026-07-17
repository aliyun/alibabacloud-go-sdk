// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeEngineJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeEngineJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeEngineJobDetailResponseBody) *GetComputeEngineJobDetailResponse
	GetBody() *GetComputeEngineJobDetailResponseBody
}

type GetComputeEngineJobDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeEngineJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeEngineJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeEngineJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeEngineJobDetailResponse) GetBody() *GetComputeEngineJobDetailResponseBody {
	return s.Body
}

func (s *GetComputeEngineJobDetailResponse) SetHeaders(v map[string]*string) *GetComputeEngineJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetComputeEngineJobDetailResponse) SetStatusCode(v int32) *GetComputeEngineJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeEngineJobDetailResponse) SetBody(v *GetComputeEngineJobDetailResponseBody) *GetComputeEngineJobDetailResponse {
	s.Body = v
	return s
}

func (s *GetComputeEngineJobDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
