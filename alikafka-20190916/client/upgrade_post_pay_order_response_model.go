// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePostPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePostPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePostPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePostPayOrderResponseBody) *UpgradePostPayOrderResponse
	GetBody() *UpgradePostPayOrderResponseBody
}

type UpgradePostPayOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePostPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePostPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePostPayOrderResponse) GetBody() *UpgradePostPayOrderResponseBody {
	return s.Body
}

func (s *UpgradePostPayOrderResponse) SetHeaders(v map[string]*string) *UpgradePostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *UpgradePostPayOrderResponse) SetStatusCode(v int32) *UpgradePostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePostPayOrderResponse) SetBody(v *UpgradePostPayOrderResponseBody) *UpgradePostPayOrderResponse {
	s.Body = v
	return s
}

func (s *UpgradePostPayOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
