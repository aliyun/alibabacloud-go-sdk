// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeContainerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeContainerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeContainerResponse
	GetStatusCode() *int32
	SetBody(v *InvokeContainerResponseBody) *InvokeContainerResponse
	GetBody() *InvokeContainerResponseBody
}

type InvokeContainerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeContainerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeContainerResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerResponse) GoString() string {
	return s.String()
}

func (s *InvokeContainerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeContainerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeContainerResponse) GetBody() *InvokeContainerResponseBody {
	return s.Body
}

func (s *InvokeContainerResponse) SetHeaders(v map[string]*string) *InvokeContainerResponse {
	s.Headers = v
	return s
}

func (s *InvokeContainerResponse) SetStatusCode(v int32) *InvokeContainerResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeContainerResponse) SetBody(v *InvokeContainerResponseBody) *InvokeContainerResponse {
	s.Body = v
	return s
}

func (s *InvokeContainerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
