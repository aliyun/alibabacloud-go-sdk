// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceControlResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceControlResponseBody) *ModifyResourceControlResponse
	GetBody() *ModifyResourceControlResponseBody
}

type ModifyResourceControlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceControlResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceControlResponse) GetBody() *ModifyResourceControlResponseBody {
	return s.Body
}

func (s *ModifyResourceControlResponse) SetHeaders(v map[string]*string) *ModifyResourceControlResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceControlResponse) SetStatusCode(v int32) *ModifyResourceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceControlResponse) SetBody(v *ModifyResourceControlResponseBody) *ModifyResourceControlResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
