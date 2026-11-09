// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterDowngradeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisasterDowngradeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisasterDowngradeResponse
	GetStatusCode() *int32
	SetBody(v *DisasterDowngradeResponseBody) *DisasterDowngradeResponse
	GetBody() *DisasterDowngradeResponseBody
}

type DisasterDowngradeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisasterDowngradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisasterDowngradeResponse) String() string {
	return dara.Prettify(s)
}

func (s DisasterDowngradeResponse) GoString() string {
	return s.String()
}

func (s *DisasterDowngradeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisasterDowngradeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisasterDowngradeResponse) GetBody() *DisasterDowngradeResponseBody {
	return s.Body
}

func (s *DisasterDowngradeResponse) SetHeaders(v map[string]*string) *DisasterDowngradeResponse {
	s.Headers = v
	return s
}

func (s *DisasterDowngradeResponse) SetStatusCode(v int32) *DisasterDowngradeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisasterDowngradeResponse) SetBody(v *DisasterDowngradeResponseBody) *DisasterDowngradeResponse {
	s.Body = v
	return s
}

func (s *DisasterDowngradeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
