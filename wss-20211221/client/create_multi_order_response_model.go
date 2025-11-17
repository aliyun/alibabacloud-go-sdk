// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiOrderResponseBody) *CreateMultiOrderResponse
	GetBody() *CreateMultiOrderResponseBody
}

type CreateMultiOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiOrderResponse) GetBody() *CreateMultiOrderResponseBody {
	return s.Body
}

func (s *CreateMultiOrderResponse) SetHeaders(v map[string]*string) *CreateMultiOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiOrderResponse) SetStatusCode(v int32) *CreateMultiOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiOrderResponse) SetBody(v *CreateMultiOrderResponseBody) *CreateMultiOrderResponse {
	s.Body = v
	return s
}

func (s *CreateMultiOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
