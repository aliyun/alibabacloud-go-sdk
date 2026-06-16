// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDisposalToolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDisposalToolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDisposalToolServiceResponse
	GetStatusCode() *int32
	SetBody(v *StartDisposalToolServiceResponseBody) *StartDisposalToolServiceResponse
	GetBody() *StartDisposalToolServiceResponseBody
}

type StartDisposalToolServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDisposalToolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDisposalToolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDisposalToolServiceResponse) GoString() string {
	return s.String()
}

func (s *StartDisposalToolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDisposalToolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDisposalToolServiceResponse) GetBody() *StartDisposalToolServiceResponseBody {
	return s.Body
}

func (s *StartDisposalToolServiceResponse) SetHeaders(v map[string]*string) *StartDisposalToolServiceResponse {
	s.Headers = v
	return s
}

func (s *StartDisposalToolServiceResponse) SetStatusCode(v int32) *StartDisposalToolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDisposalToolServiceResponse) SetBody(v *StartDisposalToolServiceResponseBody) *StartDisposalToolServiceResponse {
	s.Body = v
	return s
}

func (s *StartDisposalToolServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
