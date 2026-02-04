// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceForDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceForDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceForDeleteResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceForDeleteResponseBody) *CheckInstanceForDeleteResponse
	GetBody() *CheckInstanceForDeleteResponseBody
}

type CheckInstanceForDeleteResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceForDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceForDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceForDeleteResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceForDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceForDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceForDeleteResponse) GetBody() *CheckInstanceForDeleteResponseBody {
	return s.Body
}

func (s *CheckInstanceForDeleteResponse) SetHeaders(v map[string]*string) *CheckInstanceForDeleteResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceForDeleteResponse) SetStatusCode(v int32) *CheckInstanceForDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceForDeleteResponse) SetBody(v *CheckInstanceForDeleteResponseBody) *CheckInstanceForDeleteResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceForDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
