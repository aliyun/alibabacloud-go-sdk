// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCronClearOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataCronClearOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataCronClearOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataCronClearOrderResponseBody) *CreateDataCronClearOrderResponse
	GetBody() *CreateDataCronClearOrderResponseBody
}

type CreateDataCronClearOrderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataCronClearOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataCronClearOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataCronClearOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataCronClearOrderResponse) GetBody() *CreateDataCronClearOrderResponseBody {
	return s.Body
}

func (s *CreateDataCronClearOrderResponse) SetHeaders(v map[string]*string) *CreateDataCronClearOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCronClearOrderResponse) SetStatusCode(v int32) *CreateDataCronClearOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataCronClearOrderResponse) SetBody(v *CreateDataCronClearOrderResponseBody) *CreateDataCronClearOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataCronClearOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
