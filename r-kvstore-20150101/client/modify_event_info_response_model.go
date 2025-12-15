// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEventInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEventInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEventInfoResponseBody) *ModifyEventInfoResponse
	GetBody() *ModifyEventInfoResponseBody
}

type ModifyEventInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEventInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEventInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEventInfoResponse) GetBody() *ModifyEventInfoResponseBody {
	return s.Body
}

func (s *ModifyEventInfoResponse) SetHeaders(v map[string]*string) *ModifyEventInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventInfoResponse) SetStatusCode(v int32) *ModifyEventInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventInfoResponse) SetBody(v *ModifyEventInfoResponseBody) *ModifyEventInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyEventInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
