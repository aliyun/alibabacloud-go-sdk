// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePrePayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePrePayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePrePayOrderResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePrePayOrderResponseBody) *UpgradePrePayOrderResponse
	GetBody() *UpgradePrePayOrderResponseBody
}

type UpgradePrePayOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePrePayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderResponse) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePrePayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePrePayOrderResponse) GetBody() *UpgradePrePayOrderResponseBody {
	return s.Body
}

func (s *UpgradePrePayOrderResponse) SetHeaders(v map[string]*string) *UpgradePrePayOrderResponse {
	s.Headers = v
	return s
}

func (s *UpgradePrePayOrderResponse) SetStatusCode(v int32) *UpgradePrePayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePrePayOrderResponse) SetBody(v *UpgradePrePayOrderResponseBody) *UpgradePrePayOrderResponse {
	s.Body = v
	return s
}

func (s *UpgradePrePayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
