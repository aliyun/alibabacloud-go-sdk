// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformEipSegmentToPublicIpAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformEipSegmentToPublicIpAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformEipSegmentToPublicIpAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *TransformEipSegmentToPublicIpAddressPoolResponseBody) *TransformEipSegmentToPublicIpAddressPoolResponse
	GetBody() *TransformEipSegmentToPublicIpAddressPoolResponseBody
}

type TransformEipSegmentToPublicIpAddressPoolResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformEipSegmentToPublicIpAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformEipSegmentToPublicIpAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformEipSegmentToPublicIpAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) GetBody() *TransformEipSegmentToPublicIpAddressPoolResponseBody {
	return s.Body
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) SetHeaders(v map[string]*string) *TransformEipSegmentToPublicIpAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) SetStatusCode(v int32) *TransformEipSegmentToPublicIpAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) SetBody(v *TransformEipSegmentToPublicIpAddressPoolResponseBody) *TransformEipSegmentToPublicIpAddressPoolResponse {
	s.Body = v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}
