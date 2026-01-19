// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceServerCustomSubjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableResourceServerCustomSubjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableResourceServerCustomSubjectResponse
	GetStatusCode() *int32
	SetBody(v *DisableResourceServerCustomSubjectResponseBody) *DisableResourceServerCustomSubjectResponse
	GetBody() *DisableResourceServerCustomSubjectResponseBody
}

type DisableResourceServerCustomSubjectResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableResourceServerCustomSubjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableResourceServerCustomSubjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceServerCustomSubjectResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceServerCustomSubjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableResourceServerCustomSubjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableResourceServerCustomSubjectResponse) GetBody() *DisableResourceServerCustomSubjectResponseBody {
	return s.Body
}

func (s *DisableResourceServerCustomSubjectResponse) SetHeaders(v map[string]*string) *DisableResourceServerCustomSubjectResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceServerCustomSubjectResponse) SetStatusCode(v int32) *DisableResourceServerCustomSubjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceServerCustomSubjectResponse) SetBody(v *DisableResourceServerCustomSubjectResponseBody) *DisableResourceServerCustomSubjectResponse {
	s.Body = v
	return s
}

func (s *DisableResourceServerCustomSubjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
