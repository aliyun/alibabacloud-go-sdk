// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadFlowVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadFlowVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadFlowVersionResponse
	GetStatusCode() *int32
	SetBody(v *ReadFlowVersionResponseBody) *ReadFlowVersionResponse
	GetBody() *ReadFlowVersionResponseBody
}

type ReadFlowVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadFlowVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *ReadFlowVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadFlowVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadFlowVersionResponse) GetBody() *ReadFlowVersionResponseBody {
	return s.Body
}

func (s *ReadFlowVersionResponse) SetHeaders(v map[string]*string) *ReadFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *ReadFlowVersionResponse) SetStatusCode(v int32) *ReadFlowVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadFlowVersionResponse) SetBody(v *ReadFlowVersionResponseBody) *ReadFlowVersionResponse {
	s.Body = v
	return s
}

func (s *ReadFlowVersionResponse) Validate() error {
	return dara.Validate(s)
}
