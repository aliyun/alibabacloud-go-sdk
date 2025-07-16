// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCronScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceCronScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceCronScalerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceCronScalerResponseBody) *DeleteServiceCronScalerResponse
	GetBody() *DeleteServiceCronScalerResponseBody
}

type DeleteServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceCronScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceCronScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceCronScalerResponse) GetBody() *DeleteServiceCronScalerResponseBody {
	return s.Body
}

func (s *DeleteServiceCronScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetStatusCode(v int32) *DeleteServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetBody(v *DeleteServiceCronScalerResponseBody) *DeleteServiceCronScalerResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceCronScalerResponse) Validate() error {
	return dara.Validate(s)
}
