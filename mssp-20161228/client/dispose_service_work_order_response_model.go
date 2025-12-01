// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeServiceWorkOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisposeServiceWorkOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisposeServiceWorkOrderResponse
	GetStatusCode() *int32
	SetBody(v *DisposeServiceWorkOrderResponseBody) *DisposeServiceWorkOrderResponse
	GetBody() *DisposeServiceWorkOrderResponseBody
}

type DisposeServiceWorkOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeServiceWorkOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DisposeServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisposeServiceWorkOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisposeServiceWorkOrderResponse) GetBody() *DisposeServiceWorkOrderResponseBody {
	return s.Body
}

func (s *DisposeServiceWorkOrderResponse) SetHeaders(v map[string]*string) *DisposeServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetStatusCode(v int32) *DisposeServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetBody(v *DisposeServiceWorkOrderResponseBody) *DisposeServiceWorkOrderResponse {
	s.Body = v
	return s
}

func (s *DisposeServiceWorkOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
