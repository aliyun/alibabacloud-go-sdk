// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductEndpointsResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *GetProductEndpointsResponse
	GetBody() map[string]interface{}
}

type GetProductEndpointsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductEndpointsResponse) GoString() string {
	return s.String()
}

func (s *GetProductEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductEndpointsResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetProductEndpointsResponse) SetHeaders(v map[string]*string) *GetProductEndpointsResponse {
	s.Headers = v
	return s
}

func (s *GetProductEndpointsResponse) SetStatusCode(v int32) *GetProductEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductEndpointsResponse) SetBody(v map[string]interface{}) *GetProductEndpointsResponse {
	s.Body = v
	return s
}

func (s *GetProductEndpointsResponse) Validate() error {
	return dara.Validate(s)
}
