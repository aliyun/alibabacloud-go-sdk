// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *GetLivyComputeResponseBody) *GetLivyComputeResponse
	GetBody() *GetLivyComputeResponseBody
}

type GetLivyComputeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLivyComputeResponse) GetBody() *GetLivyComputeResponseBody {
	return s.Body
}

func (s *GetLivyComputeResponse) SetHeaders(v map[string]*string) *GetLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *GetLivyComputeResponse) SetStatusCode(v int32) *GetLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivyComputeResponse) SetBody(v *GetLivyComputeResponseBody) *GetLivyComputeResponse {
	s.Body = v
	return s
}

func (s *GetLivyComputeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
