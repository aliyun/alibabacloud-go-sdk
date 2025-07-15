// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionOccupancyOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePhysicalConnectionOccupancyOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePhysicalConnectionOccupancyOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreatePhysicalConnectionOccupancyOrderResponseBody) *CreatePhysicalConnectionOccupancyOrderResponse
	GetBody() *CreatePhysicalConnectionOccupancyOrderResponseBody
}

type CreatePhysicalConnectionOccupancyOrderResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhysicalConnectionOccupancyOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhysicalConnectionOccupancyOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionOccupancyOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) GetBody() *CreatePhysicalConnectionOccupancyOrderResponseBody {
	return s.Body
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) SetHeaders(v map[string]*string) *CreatePhysicalConnectionOccupancyOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) SetStatusCode(v int32) *CreatePhysicalConnectionOccupancyOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) SetBody(v *CreatePhysicalConnectionOccupancyOrderResponseBody) *CreatePhysicalConnectionOccupancyOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePhysicalConnectionOccupancyOrderResponse) Validate() error {
	return dara.Validate(s)
}
