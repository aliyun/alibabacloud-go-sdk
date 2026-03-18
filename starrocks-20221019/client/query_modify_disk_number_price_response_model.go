// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskNumberPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyDiskNumberPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyDiskNumberPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyDiskNumberPriceResponseBody) *QueryModifyDiskNumberPriceResponse
	GetBody() *QueryModifyDiskNumberPriceResponseBody
}

type QueryModifyDiskNumberPriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyDiskNumberPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyDiskNumberPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyDiskNumberPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyDiskNumberPriceResponse) GetBody() *QueryModifyDiskNumberPriceResponseBody {
	return s.Body
}

func (s *QueryModifyDiskNumberPriceResponse) SetHeaders(v map[string]*string) *QueryModifyDiskNumberPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponse) SetStatusCode(v int32) *QueryModifyDiskNumberPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyDiskNumberPriceResponse) SetBody(v *QueryModifyDiskNumberPriceResponseBody) *QueryModifyDiskNumberPriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyDiskNumberPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
