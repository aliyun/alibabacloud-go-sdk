// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewJobOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NewJobOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NewJobOrderResponse
	GetStatusCode() *int32
	SetBody(v *NewJobOrderResponseBody) *NewJobOrderResponse
	GetBody() *NewJobOrderResponseBody
}

type NewJobOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NewJobOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NewJobOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s NewJobOrderResponse) GoString() string {
	return s.String()
}

func (s *NewJobOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NewJobOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NewJobOrderResponse) GetBody() *NewJobOrderResponseBody {
	return s.Body
}

func (s *NewJobOrderResponse) SetHeaders(v map[string]*string) *NewJobOrderResponse {
	s.Headers = v
	return s
}

func (s *NewJobOrderResponse) SetStatusCode(v int32) *NewJobOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *NewJobOrderResponse) SetBody(v *NewJobOrderResponseBody) *NewJobOrderResponse {
	s.Body = v
	return s
}

func (s *NewJobOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
