// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLimitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeLimitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeLimitsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeLimitsResponseBody) *UpgradeLimitsResponse
	GetBody() *UpgradeLimitsResponseBody
}

type UpgradeLimitsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeLimitsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLimitsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLimitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeLimitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeLimitsResponse) GetBody() *UpgradeLimitsResponseBody {
	return s.Body
}

func (s *UpgradeLimitsResponse) SetHeaders(v map[string]*string) *UpgradeLimitsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLimitsResponse) SetStatusCode(v int32) *UpgradeLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeLimitsResponse) SetBody(v *UpgradeLimitsResponseBody) *UpgradeLimitsResponse {
	s.Body = v
	return s
}

func (s *UpgradeLimitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
