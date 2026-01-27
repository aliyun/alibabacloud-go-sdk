// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSupportDBTableRecoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SupportDBTableRecoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SupportDBTableRecoveryResponse
	GetStatusCode() *int32
	SetBody(v *SupportDBTableRecoveryResponseBody) *SupportDBTableRecoveryResponse
	GetBody() *SupportDBTableRecoveryResponseBody
}

type SupportDBTableRecoveryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SupportDBTableRecoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SupportDBTableRecoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s SupportDBTableRecoveryResponse) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SupportDBTableRecoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SupportDBTableRecoveryResponse) GetBody() *SupportDBTableRecoveryResponseBody {
	return s.Body
}

func (s *SupportDBTableRecoveryResponse) SetHeaders(v map[string]*string) *SupportDBTableRecoveryResponse {
	s.Headers = v
	return s
}

func (s *SupportDBTableRecoveryResponse) SetStatusCode(v int32) *SupportDBTableRecoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *SupportDBTableRecoveryResponse) SetBody(v *SupportDBTableRecoveryResponseBody) *SupportDBTableRecoveryResponse {
	s.Body = v
	return s
}

func (s *SupportDBTableRecoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
