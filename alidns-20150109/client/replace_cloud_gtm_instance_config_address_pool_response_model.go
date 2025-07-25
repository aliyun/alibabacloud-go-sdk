// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmInstanceConfigAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceCloudGtmInstanceConfigAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceCloudGtmInstanceConfigAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) *ReplaceCloudGtmInstanceConfigAddressPoolResponse
	GetBody() *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody
}

type ReplaceCloudGtmInstanceConfigAddressPoolResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) GetBody() *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody {
	return s.Body
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) SetHeaders(v map[string]*string) *ReplaceCloudGtmInstanceConfigAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) SetStatusCode(v int32) *ReplaceCloudGtmInstanceConfigAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) SetBody(v *ReplaceCloudGtmInstanceConfigAddressPoolResponseBody) *ReplaceCloudGtmInstanceConfigAddressPoolResponse {
	s.Body = v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}
