// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceLogStatusResponseBody) *ModifyResourceLogStatusResponse
	GetBody() *ModifyResourceLogStatusResponseBody
}

type ModifyResourceLogStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceLogStatusResponse) GetBody() *ModifyResourceLogStatusResponseBody {
	return s.Body
}

func (s *ModifyResourceLogStatusResponse) SetHeaders(v map[string]*string) *ModifyResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetStatusCode(v int32) *ModifyResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetBody(v *ModifyResourceLogStatusResponseBody) *ModifyResourceLogStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
