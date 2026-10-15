// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePriceTagDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MerchandisePriceTagDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MerchandisePriceTagDetectionResponse
	GetStatusCode() *int32
	SetBody(v *MerchandisePriceTagDetectionResponseBody) *MerchandisePriceTagDetectionResponse
	GetBody() *MerchandisePriceTagDetectionResponseBody
}

type MerchandisePriceTagDetectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MerchandisePriceTagDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MerchandisePriceTagDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionResponse) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MerchandisePriceTagDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MerchandisePriceTagDetectionResponse) GetBody() *MerchandisePriceTagDetectionResponseBody {
	return s.Body
}

func (s *MerchandisePriceTagDetectionResponse) SetHeaders(v map[string]*string) *MerchandisePriceTagDetectionResponse {
	s.Headers = v
	return s
}

func (s *MerchandisePriceTagDetectionResponse) SetStatusCode(v int32) *MerchandisePriceTagDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *MerchandisePriceTagDetectionResponse) SetBody(v *MerchandisePriceTagDetectionResponseBody) *MerchandisePriceTagDetectionResponse {
	s.Body = v
	return s
}

func (s *MerchandisePriceTagDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
