// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverCallInConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverCallInConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverCallInConfigResponse
	GetStatusCode() *int32
	SetBody(v *RecoverCallInConfigResponseBody) *RecoverCallInConfigResponse
	GetBody() *RecoverCallInConfigResponseBody
}

type RecoverCallInConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverCallInConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverCallInConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverCallInConfigResponse) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverCallInConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverCallInConfigResponse) GetBody() *RecoverCallInConfigResponseBody {
	return s.Body
}

func (s *RecoverCallInConfigResponse) SetHeaders(v map[string]*string) *RecoverCallInConfigResponse {
	s.Headers = v
	return s
}

func (s *RecoverCallInConfigResponse) SetStatusCode(v int32) *RecoverCallInConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverCallInConfigResponse) SetBody(v *RecoverCallInConfigResponseBody) *RecoverCallInConfigResponse {
	s.Body = v
	return s
}

func (s *RecoverCallInConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
