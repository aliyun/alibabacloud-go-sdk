// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDirectoryResponseBody) *ModifyDirectoryResponse
	GetBody() *ModifyDirectoryResponseBody
}

type ModifyDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDirectoryResponse) GetBody() *ModifyDirectoryResponseBody {
	return s.Body
}

func (s *ModifyDirectoryResponse) SetHeaders(v map[string]*string) *ModifyDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyDirectoryResponse) SetStatusCode(v int32) *ModifyDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDirectoryResponse) SetBody(v *ModifyDirectoryResponseBody) *ModifyDirectoryResponse {
	s.Body = v
	return s
}

func (s *ModifyDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
