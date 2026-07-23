// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseDDoSInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseDDoSInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseDDoSInstanceResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseDDoSInstanceResponseBody) *PurchaseDDoSInstanceResponse
	GetBody() *PurchaseDDoSInstanceResponseBody
}

type PurchaseDDoSInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseDDoSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseDDoSInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseDDoSInstanceResponse) GoString() string {
	return s.String()
}

func (s *PurchaseDDoSInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseDDoSInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseDDoSInstanceResponse) GetBody() *PurchaseDDoSInstanceResponseBody {
	return s.Body
}

func (s *PurchaseDDoSInstanceResponse) SetHeaders(v map[string]*string) *PurchaseDDoSInstanceResponse {
	s.Headers = v
	return s
}

func (s *PurchaseDDoSInstanceResponse) SetStatusCode(v int32) *PurchaseDDoSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseDDoSInstanceResponse) SetBody(v *PurchaseDDoSInstanceResponseBody) *PurchaseDDoSInstanceResponse {
	s.Body = v
	return s
}

func (s *PurchaseDDoSInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
