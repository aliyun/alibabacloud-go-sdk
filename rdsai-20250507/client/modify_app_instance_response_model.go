// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppInstanceResponseBody) *ModifyAppInstanceResponse
	GetBody() *ModifyAppInstanceResponseBody
}

type ModifyAppInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppInstanceResponse) GetBody() *ModifyAppInstanceResponseBody {
	return s.Body
}

func (s *ModifyAppInstanceResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceResponse) SetStatusCode(v int32) *ModifyAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceResponse) SetBody(v *ModifyAppInstanceResponseBody) *ModifyAppInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
