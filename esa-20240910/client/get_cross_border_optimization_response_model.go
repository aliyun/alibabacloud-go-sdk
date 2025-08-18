// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossBorderOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCrossBorderOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCrossBorderOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *GetCrossBorderOptimizationResponseBody) *GetCrossBorderOptimizationResponse
	GetBody() *GetCrossBorderOptimizationResponseBody
}

type GetCrossBorderOptimizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrossBorderOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrossBorderOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCrossBorderOptimizationResponse) GoString() string {
	return s.String()
}

func (s *GetCrossBorderOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCrossBorderOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCrossBorderOptimizationResponse) GetBody() *GetCrossBorderOptimizationResponseBody {
	return s.Body
}

func (s *GetCrossBorderOptimizationResponse) SetHeaders(v map[string]*string) *GetCrossBorderOptimizationResponse {
	s.Headers = v
	return s
}

func (s *GetCrossBorderOptimizationResponse) SetStatusCode(v int32) *GetCrossBorderOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrossBorderOptimizationResponse) SetBody(v *GetCrossBorderOptimizationResponseBody) *GetCrossBorderOptimizationResponse {
	s.Body = v
	return s
}

func (s *GetCrossBorderOptimizationResponse) Validate() error {
	return dara.Validate(s)
}
