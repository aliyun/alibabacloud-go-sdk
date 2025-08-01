// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListenersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceListenersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceListenersResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceListenersResponseBody) *GetServiceListenersResponse
	GetBody() *GetServiceListenersResponseBody
}

type GetServiceListenersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceListenersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceListenersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListenersResponse) GoString() string {
	return s.String()
}

func (s *GetServiceListenersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceListenersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceListenersResponse) GetBody() *GetServiceListenersResponseBody {
	return s.Body
}

func (s *GetServiceListenersResponse) SetHeaders(v map[string]*string) *GetServiceListenersResponse {
	s.Headers = v
	return s
}

func (s *GetServiceListenersResponse) SetStatusCode(v int32) *GetServiceListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceListenersResponse) SetBody(v *GetServiceListenersResponseBody) *GetServiceListenersResponse {
	s.Body = v
	return s
}

func (s *GetServiceListenersResponse) Validate() error {
	return dara.Validate(s)
}
