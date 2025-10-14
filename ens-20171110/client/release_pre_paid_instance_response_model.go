// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePrePaidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePrePaidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePrePaidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePrePaidInstanceResponseBody) *ReleasePrePaidInstanceResponse
	GetBody() *ReleasePrePaidInstanceResponseBody
}

type ReleasePrePaidInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePrePaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePrePaidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePrePaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePrePaidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePrePaidInstanceResponse) GetBody() *ReleasePrePaidInstanceResponseBody {
	return s.Body
}

func (s *ReleasePrePaidInstanceResponse) SetHeaders(v map[string]*string) *ReleasePrePaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePrePaidInstanceResponse) SetStatusCode(v int32) *ReleasePrePaidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePrePaidInstanceResponse) SetBody(v *ReleasePrePaidInstanceResponseBody) *ReleasePrePaidInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleasePrePaidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
