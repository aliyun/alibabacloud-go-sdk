// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVendorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVendorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVendorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVendorResponseBody) *DeleteVendorResponse
	GetBody() *DeleteVendorResponseBody
}

type DeleteVendorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVendorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVendorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVendorResponse) GoString() string {
	return s.String()
}

func (s *DeleteVendorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVendorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVendorResponse) GetBody() *DeleteVendorResponseBody {
	return s.Body
}

func (s *DeleteVendorResponse) SetHeaders(v map[string]*string) *DeleteVendorResponse {
	s.Headers = v
	return s
}

func (s *DeleteVendorResponse) SetStatusCode(v int32) *DeleteVendorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVendorResponse) SetBody(v *DeleteVendorResponseBody) *DeleteVendorResponse {
	s.Body = v
	return s
}

func (s *DeleteVendorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
