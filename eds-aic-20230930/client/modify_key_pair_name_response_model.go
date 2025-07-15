// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyKeyPairNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyKeyPairNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyKeyPairNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyKeyPairNameResponseBody) *ModifyKeyPairNameResponse
	GetBody() *ModifyKeyPairNameResponseBody
}

type ModifyKeyPairNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyKeyPairNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyKeyPairNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyKeyPairNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyKeyPairNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyKeyPairNameResponse) GetBody() *ModifyKeyPairNameResponseBody {
	return s.Body
}

func (s *ModifyKeyPairNameResponse) SetHeaders(v map[string]*string) *ModifyKeyPairNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyKeyPairNameResponse) SetStatusCode(v int32) *ModifyKeyPairNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyKeyPairNameResponse) SetBody(v *ModifyKeyPairNameResponseBody) *ModifyKeyPairNameResponse {
	s.Body = v
	return s
}

func (s *ModifyKeyPairNameResponse) Validate() error {
	return dara.Validate(s)
}
