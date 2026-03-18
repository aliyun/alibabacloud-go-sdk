// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifySpecTypePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifySpecTypePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifySpecTypePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifySpecTypePriceResponseBody) *QueryModifySpecTypePriceResponse
	GetBody() *QueryModifySpecTypePriceResponseBody
}

type QueryModifySpecTypePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifySpecTypePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifySpecTypePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifySpecTypePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifySpecTypePriceResponse) GetBody() *QueryModifySpecTypePriceResponseBody {
	return s.Body
}

func (s *QueryModifySpecTypePriceResponse) SetHeaders(v map[string]*string) *QueryModifySpecTypePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifySpecTypePriceResponse) SetStatusCode(v int32) *QueryModifySpecTypePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifySpecTypePriceResponse) SetBody(v *QueryModifySpecTypePriceResponseBody) *QueryModifySpecTypePriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifySpecTypePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
