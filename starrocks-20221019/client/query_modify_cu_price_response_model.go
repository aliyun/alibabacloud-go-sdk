// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyCuPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyCuPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyCuPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyCuPriceResponseBody) *QueryModifyCuPriceResponse
	GetBody() *QueryModifyCuPriceResponseBody
}

type QueryModifyCuPriceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyCuPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyCuPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyCuPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyCuPriceResponse) GetBody() *QueryModifyCuPriceResponseBody {
	return s.Body
}

func (s *QueryModifyCuPriceResponse) SetHeaders(v map[string]*string) *QueryModifyCuPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyCuPriceResponse) SetStatusCode(v int32) *QueryModifyCuPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyCuPriceResponse) SetBody(v *QueryModifyCuPriceResponseBody) *QueryModifyCuPriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyCuPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
