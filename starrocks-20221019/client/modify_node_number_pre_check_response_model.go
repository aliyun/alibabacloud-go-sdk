// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodeNumberPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodeNumberPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodeNumberPreCheckResponseBody) *ModifyNodeNumberPreCheckResponse
	GetBody() *ModifyNodeNumberPreCheckResponseBody
}

type ModifyNodeNumberPreCheckResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeNumberPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeNumberPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberPreCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodeNumberPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodeNumberPreCheckResponse) GetBody() *ModifyNodeNumberPreCheckResponseBody {
	return s.Body
}

func (s *ModifyNodeNumberPreCheckResponse) SetHeaders(v map[string]*string) *ModifyNodeNumberPreCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeNumberPreCheckResponse) SetStatusCode(v int32) *ModifyNodeNumberPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeNumberPreCheckResponse) SetBody(v *ModifyNodeNumberPreCheckResponseBody) *ModifyNodeNumberPreCheckResponse {
	s.Body = v
	return s
}

func (s *ModifyNodeNumberPreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
