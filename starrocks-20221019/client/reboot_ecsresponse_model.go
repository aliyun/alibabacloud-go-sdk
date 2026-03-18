// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootECSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootECSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootECSResponse
	GetStatusCode() *int32
	SetBody(v *RebootECSResponseBody) *RebootECSResponse
	GetBody() *RebootECSResponseBody
}

type RebootECSResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootECSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootECSResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootECSResponse) GoString() string {
	return s.String()
}

func (s *RebootECSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootECSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootECSResponse) GetBody() *RebootECSResponseBody {
	return s.Body
}

func (s *RebootECSResponse) SetHeaders(v map[string]*string) *RebootECSResponse {
	s.Headers = v
	return s
}

func (s *RebootECSResponse) SetStatusCode(v int32) *RebootECSResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootECSResponse) SetBody(v *RebootECSResponseBody) *RebootECSResponse {
	s.Body = v
	return s
}

func (s *RebootECSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
