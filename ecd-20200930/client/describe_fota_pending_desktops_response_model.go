// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaPendingDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFotaPendingDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFotaPendingDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFotaPendingDesktopsResponseBody) *DescribeFotaPendingDesktopsResponse
	GetBody() *DescribeFotaPendingDesktopsResponseBody
}

type DescribeFotaPendingDesktopsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFotaPendingDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFotaPendingDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFotaPendingDesktopsResponse) GetBody() *DescribeFotaPendingDesktopsResponseBody {
	return s.Body
}

func (s *DescribeFotaPendingDesktopsResponse) SetHeaders(v map[string]*string) *DescribeFotaPendingDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponse) SetStatusCode(v int32) *DescribeFotaPendingDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponse) SetBody(v *DescribeFotaPendingDesktopsResponseBody) *DescribeFotaPendingDesktopsResponse {
	s.Body = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
