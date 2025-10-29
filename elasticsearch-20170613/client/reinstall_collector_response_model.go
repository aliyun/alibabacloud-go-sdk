// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinstallCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinstallCollectorResponse
	GetStatusCode() *int32
	SetBody(v *ReinstallCollectorResponseBody) *ReinstallCollectorResponse
	GetBody() *ReinstallCollectorResponseBody
}

type ReinstallCollectorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReinstallCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinstallCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinstallCollectorResponse) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinstallCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinstallCollectorResponse) GetBody() *ReinstallCollectorResponseBody {
	return s.Body
}

func (s *ReinstallCollectorResponse) SetHeaders(v map[string]*string) *ReinstallCollectorResponse {
	s.Headers = v
	return s
}

func (s *ReinstallCollectorResponse) SetStatusCode(v int32) *ReinstallCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinstallCollectorResponse) SetBody(v *ReinstallCollectorResponseBody) *ReinstallCollectorResponse {
	s.Body = v
	return s
}

func (s *ReinstallCollectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
