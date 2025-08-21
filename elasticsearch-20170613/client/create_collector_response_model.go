// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCollectorResponse
	GetStatusCode() *int32
	SetBody(v *CreateCollectorResponseBody) *CreateCollectorResponse
	GetBody() *CreateCollectorResponseBody
}

type CreateCollectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectorResponse) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCollectorResponse) GetBody() *CreateCollectorResponseBody {
	return s.Body
}

func (s *CreateCollectorResponse) SetHeaders(v map[string]*string) *CreateCollectorResponse {
	s.Headers = v
	return s
}

func (s *CreateCollectorResponse) SetStatusCode(v int32) *CreateCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCollectorResponse) SetBody(v *CreateCollectorResponseBody) *CreateCollectorResponse {
	s.Body = v
	return s
}

func (s *CreateCollectorResponse) Validate() error {
	return dara.Validate(s)
}
