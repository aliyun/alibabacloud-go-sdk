// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLayerResponse
	GetStatusCode() *int32
	SetBody(v *GetLayerResponseBody) *GetLayerResponse
	GetBody() *GetLayerResponseBody
}

type GetLayerResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLayerResponse) GoString() string {
	return s.String()
}

func (s *GetLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLayerResponse) GetBody() *GetLayerResponseBody {
	return s.Body
}

func (s *GetLayerResponse) SetHeaders(v map[string]*string) *GetLayerResponse {
	s.Headers = v
	return s
}

func (s *GetLayerResponse) SetStatusCode(v int32) *GetLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerResponse) SetBody(v *GetLayerResponseBody) *GetLayerResponse {
	s.Body = v
	return s
}

func (s *GetLayerResponse) Validate() error {
	return dara.Validate(s)
}
