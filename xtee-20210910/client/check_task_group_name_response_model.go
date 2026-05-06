// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTaskGroupNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTaskGroupNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTaskGroupNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckTaskGroupNameResponseBody) *CheckTaskGroupNameResponse
	GetBody() *CheckTaskGroupNameResponseBody
}

type CheckTaskGroupNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTaskGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTaskGroupNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTaskGroupNameResponse) GoString() string {
	return s.String()
}

func (s *CheckTaskGroupNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTaskGroupNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTaskGroupNameResponse) GetBody() *CheckTaskGroupNameResponseBody {
	return s.Body
}

func (s *CheckTaskGroupNameResponse) SetHeaders(v map[string]*string) *CheckTaskGroupNameResponse {
	s.Headers = v
	return s
}

func (s *CheckTaskGroupNameResponse) SetStatusCode(v int32) *CheckTaskGroupNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTaskGroupNameResponse) SetBody(v *CheckTaskGroupNameResponseBody) *CheckTaskGroupNameResponse {
	s.Body = v
	return s
}

func (s *CheckTaskGroupNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
