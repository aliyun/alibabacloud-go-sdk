// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuySecretNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuySecretNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuySecretNoResponse
	GetStatusCode() *int32
	SetBody(v *BuySecretNoResponseBody) *BuySecretNoResponse
	GetBody() *BuySecretNoResponseBody
}

type BuySecretNoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuySecretNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuySecretNoResponse) String() string {
	return dara.Prettify(s)
}

func (s BuySecretNoResponse) GoString() string {
	return s.String()
}

func (s *BuySecretNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuySecretNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuySecretNoResponse) GetBody() *BuySecretNoResponseBody {
	return s.Body
}

func (s *BuySecretNoResponse) SetHeaders(v map[string]*string) *BuySecretNoResponse {
	s.Headers = v
	return s
}

func (s *BuySecretNoResponse) SetStatusCode(v int32) *BuySecretNoResponse {
	s.StatusCode = &v
	return s
}

func (s *BuySecretNoResponse) SetBody(v *BuySecretNoResponseBody) *BuySecretNoResponse {
	s.Body = v
	return s
}

func (s *BuySecretNoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
