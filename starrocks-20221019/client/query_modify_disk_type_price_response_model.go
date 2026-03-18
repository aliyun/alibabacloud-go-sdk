// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskTypePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyDiskTypePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyDiskTypePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyDiskTypePriceResponseBody) *QueryModifyDiskTypePriceResponse
	GetBody() *QueryModifyDiskTypePriceResponseBody
}

type QueryModifyDiskTypePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyDiskTypePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyDiskTypePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyDiskTypePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyDiskTypePriceResponse) GetBody() *QueryModifyDiskTypePriceResponseBody {
	return s.Body
}

func (s *QueryModifyDiskTypePriceResponse) SetHeaders(v map[string]*string) *QueryModifyDiskTypePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyDiskTypePriceResponse) SetStatusCode(v int32) *QueryModifyDiskTypePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyDiskTypePriceResponse) SetBody(v *QueryModifyDiskTypePriceResponseBody) *QueryModifyDiskTypePriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyDiskTypePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
