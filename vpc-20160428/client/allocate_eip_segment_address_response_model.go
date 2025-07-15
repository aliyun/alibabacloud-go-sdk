// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipSegmentAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateEipSegmentAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateEipSegmentAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateEipSegmentAddressResponseBody) *AllocateEipSegmentAddressResponse
	GetBody() *AllocateEipSegmentAddressResponseBody
}

type AllocateEipSegmentAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateEipSegmentAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateEipSegmentAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipSegmentAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateEipSegmentAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateEipSegmentAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateEipSegmentAddressResponse) GetBody() *AllocateEipSegmentAddressResponseBody {
	return s.Body
}

func (s *AllocateEipSegmentAddressResponse) SetHeaders(v map[string]*string) *AllocateEipSegmentAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateEipSegmentAddressResponse) SetStatusCode(v int32) *AllocateEipSegmentAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateEipSegmentAddressResponse) SetBody(v *AllocateEipSegmentAddressResponseBody) *AllocateEipSegmentAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateEipSegmentAddressResponse) Validate() error {
	return dara.Validate(s)
}
