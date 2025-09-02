// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPublicModelEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPublicModelEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPublicModelEngineResponse
	GetStatusCode() *int32
	SetBody(v *QueryPublicModelEngineResponseBody) *QueryPublicModelEngineResponse
	GetBody() *QueryPublicModelEngineResponseBody
}

type QueryPublicModelEngineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPublicModelEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPublicModelEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPublicModelEngineResponse) GoString() string {
	return s.String()
}

func (s *QueryPublicModelEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPublicModelEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPublicModelEngineResponse) GetBody() *QueryPublicModelEngineResponseBody {
	return s.Body
}

func (s *QueryPublicModelEngineResponse) SetHeaders(v map[string]*string) *QueryPublicModelEngineResponse {
	s.Headers = v
	return s
}

func (s *QueryPublicModelEngineResponse) SetStatusCode(v int32) *QueryPublicModelEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPublicModelEngineResponse) SetBody(v *QueryPublicModelEngineResponseBody) *QueryPublicModelEngineResponse {
	s.Body = v
	return s
}

func (s *QueryPublicModelEngineResponse) Validate() error {
	return dara.Validate(s)
}
