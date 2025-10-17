// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppCallbackStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppCallbackStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppCallbackStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppCallbackStatusResponseBody) *ModifyAppCallbackStatusResponse
	GetBody() *ModifyAppCallbackStatusResponseBody
}

type ModifyAppCallbackStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppCallbackStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppCallbackStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppCallbackStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppCallbackStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppCallbackStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppCallbackStatusResponse) GetBody() *ModifyAppCallbackStatusResponseBody {
	return s.Body
}

func (s *ModifyAppCallbackStatusResponse) SetHeaders(v map[string]*string) *ModifyAppCallbackStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppCallbackStatusResponse) SetStatusCode(v int32) *ModifyAppCallbackStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppCallbackStatusResponse) SetBody(v *ModifyAppCallbackStatusResponseBody) *ModifyAppCallbackStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAppCallbackStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
