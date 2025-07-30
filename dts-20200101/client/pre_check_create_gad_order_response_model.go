// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateGadOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreCheckCreateGadOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreCheckCreateGadOrderResponse
	GetStatusCode() *int32
	SetBody(v *PreCheckCreateGadOrderResponseBody) *PreCheckCreateGadOrderResponse
	GetBody() *PreCheckCreateGadOrderResponseBody
}

type PreCheckCreateGadOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreCheckCreateGadOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCheckCreateGadOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateGadOrderResponse) GoString() string {
	return s.String()
}

func (s *PreCheckCreateGadOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreCheckCreateGadOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreCheckCreateGadOrderResponse) GetBody() *PreCheckCreateGadOrderResponseBody {
	return s.Body
}

func (s *PreCheckCreateGadOrderResponse) SetHeaders(v map[string]*string) *PreCheckCreateGadOrderResponse {
	s.Headers = v
	return s
}

func (s *PreCheckCreateGadOrderResponse) SetStatusCode(v int32) *PreCheckCreateGadOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCheckCreateGadOrderResponse) SetBody(v *PreCheckCreateGadOrderResponseBody) *PreCheckCreateGadOrderResponse {
	s.Body = v
	return s
}

func (s *PreCheckCreateGadOrderResponse) Validate() error {
	return dara.Validate(s)
}
