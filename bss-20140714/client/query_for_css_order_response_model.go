// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForCssOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryForCssOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryForCssOrderResponse
	GetStatusCode() *int32
	SetBody(v *QueryForCssOrderResponseBody) *QueryForCssOrderResponse
	GetBody() *QueryForCssOrderResponseBody
}

type QueryForCssOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryForCssOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryForCssOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryForCssOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryForCssOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryForCssOrderResponse) GetBody() *QueryForCssOrderResponseBody {
	return s.Body
}

func (s *QueryForCssOrderResponse) SetHeaders(v map[string]*string) *QueryForCssOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryForCssOrderResponse) SetStatusCode(v int32) *QueryForCssOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryForCssOrderResponse) SetBody(v *QueryForCssOrderResponseBody) *QueryForCssOrderResponse {
	s.Body = v
	return s
}

func (s *QueryForCssOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
