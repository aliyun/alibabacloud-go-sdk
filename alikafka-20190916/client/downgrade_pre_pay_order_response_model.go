// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePrePayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradePrePayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradePrePayOrderResponse
	GetStatusCode() *int32
	SetBody(v *DowngradePrePayOrderResponseBody) *DowngradePrePayOrderResponse
	GetBody() *DowngradePrePayOrderResponseBody
}

type DowngradePrePayOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradePrePayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradePrePayOrderResponse) GoString() string {
	return s.String()
}

func (s *DowngradePrePayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradePrePayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradePrePayOrderResponse) GetBody() *DowngradePrePayOrderResponseBody {
	return s.Body
}

func (s *DowngradePrePayOrderResponse) SetHeaders(v map[string]*string) *DowngradePrePayOrderResponse {
	s.Headers = v
	return s
}

func (s *DowngradePrePayOrderResponse) SetStatusCode(v int32) *DowngradePrePayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradePrePayOrderResponse) SetBody(v *DowngradePrePayOrderResponseBody) *DowngradePrePayOrderResponse {
	s.Body = v
	return s
}

func (s *DowngradePrePayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
