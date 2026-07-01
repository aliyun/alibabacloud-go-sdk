// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemEventAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySystemEventAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySystemEventAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySystemEventAttributeResponseBody) *ModifySystemEventAttributeResponse
	GetBody() *ModifySystemEventAttributeResponseBody
}

type ModifySystemEventAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySystemEventAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySystemEventAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemEventAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySystemEventAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySystemEventAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySystemEventAttributeResponse) GetBody() *ModifySystemEventAttributeResponseBody {
	return s.Body
}

func (s *ModifySystemEventAttributeResponse) SetHeaders(v map[string]*string) *ModifySystemEventAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySystemEventAttributeResponse) SetStatusCode(v int32) *ModifySystemEventAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySystemEventAttributeResponse) SetBody(v *ModifySystemEventAttributeResponseBody) *ModifySystemEventAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifySystemEventAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
