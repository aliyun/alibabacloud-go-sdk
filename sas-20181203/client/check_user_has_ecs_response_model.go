// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserHasEcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserHasEcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserHasEcsResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserHasEcsResponseBody) *CheckUserHasEcsResponse
	GetBody() *CheckUserHasEcsResponseBody
}

type CheckUserHasEcsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserHasEcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserHasEcsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserHasEcsResponse) GoString() string {
	return s.String()
}

func (s *CheckUserHasEcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserHasEcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserHasEcsResponse) GetBody() *CheckUserHasEcsResponseBody {
	return s.Body
}

func (s *CheckUserHasEcsResponse) SetHeaders(v map[string]*string) *CheckUserHasEcsResponse {
	s.Headers = v
	return s
}

func (s *CheckUserHasEcsResponse) SetStatusCode(v int32) *CheckUserHasEcsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserHasEcsResponse) SetBody(v *CheckUserHasEcsResponseBody) *CheckUserHasEcsResponse {
	s.Body = v
	return s
}

func (s *CheckUserHasEcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
