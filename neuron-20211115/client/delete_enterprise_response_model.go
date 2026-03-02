// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnterpriseResponse
	GetStatusCode() *int32
}

type DeleteEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnterpriseResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseResponse) SetStatusCode(v int32) *DeleteEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseResponse) Validate() error {
	return dara.Validate(s)
}
