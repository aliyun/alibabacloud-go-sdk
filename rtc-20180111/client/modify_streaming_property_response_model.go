// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStreamingPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStreamingPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStreamingPropertyResponseBody) *ModifyStreamingPropertyResponse
	GetBody() *ModifyStreamingPropertyResponseBody
}

type ModifyStreamingPropertyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStreamingPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStreamingPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyStreamingPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStreamingPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStreamingPropertyResponse) GetBody() *ModifyStreamingPropertyResponseBody {
	return s.Body
}

func (s *ModifyStreamingPropertyResponse) SetHeaders(v map[string]*string) *ModifyStreamingPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyStreamingPropertyResponse) SetStatusCode(v int32) *ModifyStreamingPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStreamingPropertyResponse) SetBody(v *ModifyStreamingPropertyResponseBody) *ModifyStreamingPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyStreamingPropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
