// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipSegmentAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseEipSegmentAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseEipSegmentAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseEipSegmentAddressResponseBody) *ReleaseEipSegmentAddressResponse
	GetBody() *ReleaseEipSegmentAddressResponseBody
}

type ReleaseEipSegmentAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseEipSegmentAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseEipSegmentAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipSegmentAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseEipSegmentAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseEipSegmentAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseEipSegmentAddressResponse) GetBody() *ReleaseEipSegmentAddressResponseBody {
	return s.Body
}

func (s *ReleaseEipSegmentAddressResponse) SetHeaders(v map[string]*string) *ReleaseEipSegmentAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseEipSegmentAddressResponse) SetStatusCode(v int32) *ReleaseEipSegmentAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseEipSegmentAddressResponse) SetBody(v *ReleaseEipSegmentAddressResponseBody) *ReleaseEipSegmentAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseEipSegmentAddressResponse) Validate() error {
	return dara.Validate(s)
}
