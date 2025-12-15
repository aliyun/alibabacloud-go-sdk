// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTaskInfoResponseBody) *ModifyTaskInfoResponse
	GetBody() *ModifyTaskInfoResponseBody
}

type ModifyTaskInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTaskInfoResponse) GetBody() *ModifyTaskInfoResponseBody {
	return s.Body
}

func (s *ModifyTaskInfoResponse) SetHeaders(v map[string]*string) *ModifyTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyTaskInfoResponse) SetStatusCode(v int32) *ModifyTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTaskInfoResponse) SetBody(v *ModifyTaskInfoResponseBody) *ModifyTaskInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
