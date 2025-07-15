// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpsItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpsItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpsItemResponseBody) *CreateOpsItemResponse
	GetBody() *CreateOpsItemResponseBody
}

type CreateOpsItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpsItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsItemResponse) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpsItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpsItemResponse) GetBody() *CreateOpsItemResponseBody {
	return s.Body
}

func (s *CreateOpsItemResponse) SetHeaders(v map[string]*string) *CreateOpsItemResponse {
	s.Headers = v
	return s
}

func (s *CreateOpsItemResponse) SetStatusCode(v int32) *CreateOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpsItemResponse) SetBody(v *CreateOpsItemResponseBody) *CreateOpsItemResponse {
	s.Body = v
	return s
}

func (s *CreateOpsItemResponse) Validate() error {
	return dara.Validate(s)
}
