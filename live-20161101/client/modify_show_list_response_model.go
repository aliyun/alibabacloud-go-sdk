// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyShowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyShowListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyShowListResponseBody) *ModifyShowListResponse
	GetBody() *ModifyShowListResponseBody
}

type ModifyShowListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyShowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyShowListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyShowListResponse) GoString() string {
	return s.String()
}

func (s *ModifyShowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyShowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyShowListResponse) GetBody() *ModifyShowListResponseBody {
	return s.Body
}

func (s *ModifyShowListResponse) SetHeaders(v map[string]*string) *ModifyShowListResponse {
	s.Headers = v
	return s
}

func (s *ModifyShowListResponse) SetStatusCode(v int32) *ModifyShowListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyShowListResponse) SetBody(v *ModifyShowListResponseBody) *ModifyShowListResponse {
	s.Body = v
	return s
}

func (s *ModifyShowListResponse) Validate() error {
	return dara.Validate(s)
}
