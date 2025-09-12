// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowLedgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowLedgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowLedgeResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowLedgeResponseBody) *CreateKnowLedgeResponse
	GetBody() *CreateKnowLedgeResponseBody
}

type CreateKnowLedgeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowLedgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowLedgeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowLedgeResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowLedgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowLedgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowLedgeResponse) GetBody() *CreateKnowLedgeResponseBody {
	return s.Body
}

func (s *CreateKnowLedgeResponse) SetHeaders(v map[string]*string) *CreateKnowLedgeResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowLedgeResponse) SetStatusCode(v int32) *CreateKnowLedgeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowLedgeResponse) SetBody(v *CreateKnowLedgeResponseBody) *CreateKnowLedgeResponse {
	s.Body = v
	return s
}

func (s *CreateKnowLedgeResponse) Validate() error {
	return dara.Validate(s)
}
