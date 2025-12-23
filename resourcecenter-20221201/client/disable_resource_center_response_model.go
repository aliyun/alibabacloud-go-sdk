// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableResourceCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableResourceCenterResponse
	GetStatusCode() *int32
	SetBody(v *DisableResourceCenterResponseBody) *DisableResourceCenterResponse
	GetBody() *DisableResourceCenterResponseBody
}

type DisableResourceCenterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableResourceCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableResourceCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableResourceCenterResponse) GetBody() *DisableResourceCenterResponseBody {
	return s.Body
}

func (s *DisableResourceCenterResponse) SetHeaders(v map[string]*string) *DisableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceCenterResponse) SetStatusCode(v int32) *DisableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceCenterResponse) SetBody(v *DisableResourceCenterResponseBody) *DisableResourceCenterResponse {
	s.Body = v
	return s
}

func (s *DisableResourceCenterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
