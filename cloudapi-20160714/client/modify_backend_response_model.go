// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackendResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackendResponseBody) *ModifyBackendResponse
	GetBody() *ModifyBackendResponseBody
}

type ModifyBackendResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackendResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackendResponse) GetBody() *ModifyBackendResponseBody {
	return s.Body
}

func (s *ModifyBackendResponse) SetHeaders(v map[string]*string) *ModifyBackendResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackendResponse) SetStatusCode(v int32) *ModifyBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackendResponse) SetBody(v *ModifyBackendResponseBody) *ModifyBackendResponse {
	s.Body = v
	return s
}

func (s *ModifyBackendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
