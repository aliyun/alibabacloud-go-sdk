// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueSoftphoneCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IssueSoftphoneCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IssueSoftphoneCommandResponse
	GetStatusCode() *int32
	SetBody(v *IssueSoftphoneCommandResponseBody) *IssueSoftphoneCommandResponse
	GetBody() *IssueSoftphoneCommandResponseBody
}

type IssueSoftphoneCommandResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IssueSoftphoneCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IssueSoftphoneCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s IssueSoftphoneCommandResponse) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IssueSoftphoneCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IssueSoftphoneCommandResponse) GetBody() *IssueSoftphoneCommandResponseBody {
	return s.Body
}

func (s *IssueSoftphoneCommandResponse) SetHeaders(v map[string]*string) *IssueSoftphoneCommandResponse {
	s.Headers = v
	return s
}

func (s *IssueSoftphoneCommandResponse) SetStatusCode(v int32) *IssueSoftphoneCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueSoftphoneCommandResponse) SetBody(v *IssueSoftphoneCommandResponseBody) *IssueSoftphoneCommandResponse {
	s.Body = v
	return s
}

func (s *IssueSoftphoneCommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
