// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRemediationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRemediationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRemediationResponse
	GetStatusCode() *int32
	SetBody(v *StartRemediationResponseBody) *StartRemediationResponse
	GetBody() *StartRemediationResponseBody
}

type StartRemediationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRemediationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRemediationResponse) GoString() string {
	return s.String()
}

func (s *StartRemediationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRemediationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRemediationResponse) GetBody() *StartRemediationResponseBody {
	return s.Body
}

func (s *StartRemediationResponse) SetHeaders(v map[string]*string) *StartRemediationResponse {
	s.Headers = v
	return s
}

func (s *StartRemediationResponse) SetStatusCode(v int32) *StartRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRemediationResponse) SetBody(v *StartRemediationResponseBody) *StartRemediationResponse {
	s.Body = v
	return s
}

func (s *StartRemediationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
