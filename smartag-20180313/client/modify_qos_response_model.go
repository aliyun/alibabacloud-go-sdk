// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQosResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQosResponseBody) *ModifyQosResponse
	GetBody() *ModifyQosResponseBody
}

type ModifyQosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQosResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosResponse) GoString() string {
	return s.String()
}

func (s *ModifyQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQosResponse) GetBody() *ModifyQosResponseBody {
	return s.Body
}

func (s *ModifyQosResponse) SetHeaders(v map[string]*string) *ModifyQosResponse {
	s.Headers = v
	return s
}

func (s *ModifyQosResponse) SetStatusCode(v int32) *ModifyQosResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQosResponse) SetBody(v *ModifyQosResponseBody) *ModifyQosResponse {
	s.Body = v
	return s
}

func (s *ModifyQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
