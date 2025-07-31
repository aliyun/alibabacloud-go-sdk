// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyServiceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyServiceInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyServiceInfoResponseBody) *ModifyServiceInfoResponse
	GetBody() *ModifyServiceInfoResponseBody
}

type ModifyServiceInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyServiceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyServiceInfoResponse) GetBody() *ModifyServiceInfoResponseBody {
	return s.Body
}

func (s *ModifyServiceInfoResponse) SetHeaders(v map[string]*string) *ModifyServiceInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceInfoResponse) SetStatusCode(v int32) *ModifyServiceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceInfoResponse) SetBody(v *ModifyServiceInfoResponseBody) *ModifyServiceInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyServiceInfoResponse) Validate() error {
	return dara.Validate(s)
}
