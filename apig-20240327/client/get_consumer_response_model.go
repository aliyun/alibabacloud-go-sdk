// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerResponseBody) *GetConsumerResponse
	GetBody() *GetConsumerResponseBody
}

type GetConsumerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerResponse) GetBody() *GetConsumerResponseBody {
	return s.Body
}

func (s *GetConsumerResponse) SetHeaders(v map[string]*string) *GetConsumerResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerResponse) SetStatusCode(v int32) *GetConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerResponse) SetBody(v *GetConsumerResponseBody) *GetConsumerResponse {
	s.Body = v
	return s
}

func (s *GetConsumerResponse) Validate() error {
	return dara.Validate(s)
}
