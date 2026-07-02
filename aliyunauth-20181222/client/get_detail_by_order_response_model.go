// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetailByOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetailByOrderResponse
	GetStatusCode() *int32
	SetBody(v *GetDetailByOrderResponseBody) *GetDetailByOrderResponse
	GetBody() *GetDetailByOrderResponseBody
}

type GetDetailByOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetailByOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetailByOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByOrderResponse) GoString() string {
	return s.String()
}

func (s *GetDetailByOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetailByOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetailByOrderResponse) GetBody() *GetDetailByOrderResponseBody {
	return s.Body
}

func (s *GetDetailByOrderResponse) SetHeaders(v map[string]*string) *GetDetailByOrderResponse {
	s.Headers = v
	return s
}

func (s *GetDetailByOrderResponse) SetStatusCode(v int32) *GetDetailByOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetailByOrderResponse) SetBody(v *GetDetailByOrderResponseBody) *GetDetailByOrderResponse {
	s.Body = v
	return s
}

func (s *GetDetailByOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
