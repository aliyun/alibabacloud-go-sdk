// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseBotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseBotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseBotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseBotInstanceResponseBody) *PurchaseBotInstanceResponse
	GetBody() *PurchaseBotInstanceResponseBody
}

type PurchaseBotInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseBotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseBotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseBotInstanceResponse) GoString() string {
	return s.String()
}

func (s *PurchaseBotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseBotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseBotInstanceResponse) GetBody() *PurchaseBotInstanceResponseBody {
	return s.Body
}

func (s *PurchaseBotInstanceResponse) SetHeaders(v map[string]*string) *PurchaseBotInstanceResponse {
	s.Headers = v
	return s
}

func (s *PurchaseBotInstanceResponse) SetStatusCode(v int32) *PurchaseBotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseBotInstanceResponse) SetBody(v *PurchaseBotInstanceResponseBody) *PurchaseBotInstanceResponse {
	s.Body = v
	return s
}

func (s *PurchaseBotInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
