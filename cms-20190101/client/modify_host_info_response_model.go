// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostInfoResponseBody) *ModifyHostInfoResponse
	GetBody() *ModifyHostInfoResponseBody
}

type ModifyHostInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostInfoResponse) GetBody() *ModifyHostInfoResponseBody {
	return s.Body
}

func (s *ModifyHostInfoResponse) SetHeaders(v map[string]*string) *ModifyHostInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostInfoResponse) SetStatusCode(v int32) *ModifyHostInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostInfoResponse) SetBody(v *ModifyHostInfoResponseBody) *ModifyHostInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyHostInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
