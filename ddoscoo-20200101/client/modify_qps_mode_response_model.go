// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQpsModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQpsModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQpsModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQpsModeResponseBody) *ModifyQpsModeResponse
	GetBody() *ModifyQpsModeResponseBody
}

type ModifyQpsModeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQpsModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQpsModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQpsModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyQpsModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQpsModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQpsModeResponse) GetBody() *ModifyQpsModeResponseBody {
	return s.Body
}

func (s *ModifyQpsModeResponse) SetHeaders(v map[string]*string) *ModifyQpsModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyQpsModeResponse) SetStatusCode(v int32) *ModifyQpsModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQpsModeResponse) SetBody(v *ModifyQpsModeResponseBody) *ModifyQpsModeResponse {
	s.Body = v
	return s
}

func (s *ModifyQpsModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
