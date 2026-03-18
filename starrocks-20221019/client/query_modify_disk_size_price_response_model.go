// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskSizePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyDiskSizePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyDiskSizePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyDiskSizePriceResponseBody) *QueryModifyDiskSizePriceResponse
	GetBody() *QueryModifyDiskSizePriceResponseBody
}

type QueryModifyDiskSizePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyDiskSizePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyDiskSizePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyDiskSizePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyDiskSizePriceResponse) GetBody() *QueryModifyDiskSizePriceResponseBody {
	return s.Body
}

func (s *QueryModifyDiskSizePriceResponse) SetHeaders(v map[string]*string) *QueryModifyDiskSizePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyDiskSizePriceResponse) SetStatusCode(v int32) *QueryModifyDiskSizePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyDiskSizePriceResponse) SetBody(v *QueryModifyDiskSizePriceResponseBody) *QueryModifyDiskSizePriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyDiskSizePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
