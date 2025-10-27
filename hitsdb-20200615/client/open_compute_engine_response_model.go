// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenComputeEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenComputeEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenComputeEngineResponse
	GetStatusCode() *int32
	SetBody(v *OpenComputeEngineResponseBody) *OpenComputeEngineResponse
	GetBody() *OpenComputeEngineResponseBody
}

type OpenComputeEngineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenComputeEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenComputeEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenComputeEngineResponse) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenComputeEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenComputeEngineResponse) GetBody() *OpenComputeEngineResponseBody {
	return s.Body
}

func (s *OpenComputeEngineResponse) SetHeaders(v map[string]*string) *OpenComputeEngineResponse {
	s.Headers = v
	return s
}

func (s *OpenComputeEngineResponse) SetStatusCode(v int32) *OpenComputeEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenComputeEngineResponse) SetBody(v *OpenComputeEngineResponseBody) *OpenComputeEngineResponse {
	s.Body = v
	return s
}

func (s *OpenComputeEngineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
