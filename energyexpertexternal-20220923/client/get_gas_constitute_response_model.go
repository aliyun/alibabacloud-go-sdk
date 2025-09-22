// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGasConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGasConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGasConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetGasConstituteResponseBody) *GetGasConstituteResponse
	GetBody() *GetGasConstituteResponseBody
}

type GetGasConstituteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGasConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGasConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGasConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGasConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGasConstituteResponse) GetBody() *GetGasConstituteResponseBody {
	return s.Body
}

func (s *GetGasConstituteResponse) SetHeaders(v map[string]*string) *GetGasConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetGasConstituteResponse) SetStatusCode(v int32) *GetGasConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGasConstituteResponse) SetBody(v *GetGasConstituteResponseBody) *GetGasConstituteResponse {
	s.Body = v
	return s
}

func (s *GetGasConstituteResponse) Validate() error {
	return dara.Validate(s)
}
