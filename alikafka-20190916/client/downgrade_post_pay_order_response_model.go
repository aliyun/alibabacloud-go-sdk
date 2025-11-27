// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePostPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradePostPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradePostPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *DowngradePostPayOrderResponseBody) *DowngradePostPayOrderResponse
	GetBody() *DowngradePostPayOrderResponseBody
}

type DowngradePostPayOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradePostPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradePostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *DowngradePostPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradePostPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradePostPayOrderResponse) GetBody() *DowngradePostPayOrderResponseBody {
	return s.Body
}

func (s *DowngradePostPayOrderResponse) SetHeaders(v map[string]*string) *DowngradePostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *DowngradePostPayOrderResponse) SetStatusCode(v int32) *DowngradePostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradePostPayOrderResponse) SetBody(v *DowngradePostPayOrderResponseBody) *DowngradePostPayOrderResponse {
	s.Body = v
	return s
}

func (s *DowngradePostPayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
