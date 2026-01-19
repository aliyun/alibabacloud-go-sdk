// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationResourceServerIdentifierResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationResourceServerIdentifierResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationResourceServerIdentifierResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationResourceServerIdentifierResponseBody) *SetApplicationResourceServerIdentifierResponse
	GetBody() *SetApplicationResourceServerIdentifierResponseBody
}

type SetApplicationResourceServerIdentifierResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationResourceServerIdentifierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationResourceServerIdentifierResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationResourceServerIdentifierResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationResourceServerIdentifierResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationResourceServerIdentifierResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationResourceServerIdentifierResponse) GetBody() *SetApplicationResourceServerIdentifierResponseBody {
	return s.Body
}

func (s *SetApplicationResourceServerIdentifierResponse) SetHeaders(v map[string]*string) *SetApplicationResourceServerIdentifierResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationResourceServerIdentifierResponse) SetStatusCode(v int32) *SetApplicationResourceServerIdentifierResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierResponse) SetBody(v *SetApplicationResourceServerIdentifierResponseBody) *SetApplicationResourceServerIdentifierResponse {
	s.Body = v
	return s
}

func (s *SetApplicationResourceServerIdentifierResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
