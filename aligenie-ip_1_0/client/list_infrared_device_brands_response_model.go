// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredDeviceBrandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInfraredDeviceBrandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInfraredDeviceBrandsResponse
	GetStatusCode() *int32
	SetBody(v *ListInfraredDeviceBrandsResponseBody) *ListInfraredDeviceBrandsResponse
	GetBody() *ListInfraredDeviceBrandsResponseBody
}

type ListInfraredDeviceBrandsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInfraredDeviceBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInfraredDeviceBrandsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredDeviceBrandsResponse) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInfraredDeviceBrandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInfraredDeviceBrandsResponse) GetBody() *ListInfraredDeviceBrandsResponseBody {
	return s.Body
}

func (s *ListInfraredDeviceBrandsResponse) SetHeaders(v map[string]*string) *ListInfraredDeviceBrandsResponse {
	s.Headers = v
	return s
}

func (s *ListInfraredDeviceBrandsResponse) SetStatusCode(v int32) *ListInfraredDeviceBrandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponse) SetBody(v *ListInfraredDeviceBrandsResponseBody) *ListInfraredDeviceBrandsResponse {
	s.Body = v
	return s
}

func (s *ListInfraredDeviceBrandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
