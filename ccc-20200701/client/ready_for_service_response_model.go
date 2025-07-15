// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadyForServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadyForServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadyForServiceResponse
	GetStatusCode() *int32
	SetBody(v *ReadyForServiceResponseBody) *ReadyForServiceResponse
	GetBody() *ReadyForServiceResponseBody
}

type ReadyForServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadyForServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadyForServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadyForServiceResponse) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadyForServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadyForServiceResponse) GetBody() *ReadyForServiceResponseBody {
	return s.Body
}

func (s *ReadyForServiceResponse) SetHeaders(v map[string]*string) *ReadyForServiceResponse {
	s.Headers = v
	return s
}

func (s *ReadyForServiceResponse) SetStatusCode(v int32) *ReadyForServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadyForServiceResponse) SetBody(v *ReadyForServiceResponseBody) *ReadyForServiceResponse {
	s.Body = v
	return s
}

func (s *ReadyForServiceResponse) Validate() error {
	return dara.Validate(s)
}
