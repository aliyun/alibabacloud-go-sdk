// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgorithmDefsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgorithmDefsResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgorithmDefsResponseBody) *GetAlgorithmDefsResponse
	GetBody() *GetAlgorithmDefsResponseBody
}

type GetAlgorithmDefsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmDefsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmDefsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefsResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgorithmDefsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgorithmDefsResponse) GetBody() *GetAlgorithmDefsResponseBody {
	return s.Body
}

func (s *GetAlgorithmDefsResponse) SetHeaders(v map[string]*string) *GetAlgorithmDefsResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDefsResponse) SetStatusCode(v int32) *GetAlgorithmDefsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmDefsResponse) SetBody(v *GetAlgorithmDefsResponseBody) *GetAlgorithmDefsResponse {
	s.Body = v
	return s
}

func (s *GetAlgorithmDefsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
