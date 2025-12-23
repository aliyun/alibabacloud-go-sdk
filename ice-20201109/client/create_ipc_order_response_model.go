// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpcOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpcOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpcOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpcOrderResponseBody) *CreateIpcOrderResponse
	GetBody() *CreateIpcOrderResponseBody
}

type CreateIpcOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpcOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpcOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpcOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateIpcOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpcOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpcOrderResponse) GetBody() *CreateIpcOrderResponseBody {
	return s.Body
}

func (s *CreateIpcOrderResponse) SetHeaders(v map[string]*string) *CreateIpcOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateIpcOrderResponse) SetStatusCode(v int32) *CreateIpcOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpcOrderResponse) SetBody(v *CreateIpcOrderResponseBody) *CreateIpcOrderResponse {
	s.Body = v
	return s
}

func (s *CreateIpcOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
