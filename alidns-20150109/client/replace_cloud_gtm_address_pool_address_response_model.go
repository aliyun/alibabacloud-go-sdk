// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmAddressPoolAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceCloudGtmAddressPoolAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceCloudGtmAddressPoolAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceCloudGtmAddressPoolAddressResponseBody) *ReplaceCloudGtmAddressPoolAddressResponse
	GetBody() *ReplaceCloudGtmAddressPoolAddressResponseBody
}

type ReplaceCloudGtmAddressPoolAddressResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceCloudGtmAddressPoolAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceCloudGtmAddressPoolAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmAddressPoolAddressResponse) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) GetBody() *ReplaceCloudGtmAddressPoolAddressResponseBody {
	return s.Body
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) SetHeaders(v map[string]*string) *ReplaceCloudGtmAddressPoolAddressResponse {
	s.Headers = v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) SetStatusCode(v int32) *ReplaceCloudGtmAddressPoolAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) SetBody(v *ReplaceCloudGtmAddressPoolAddressResponseBody) *ReplaceCloudGtmAddressPoolAddressResponse {
	s.Body = v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
