// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpsItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpsItemResponse
	GetStatusCode() *int32
	SetBody(v *GetOpsItemResponseBody) *GetOpsItemResponse
	GetBody() *GetOpsItemResponseBody
}

type GetOpsItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpsItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpsItemResponse) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpsItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpsItemResponse) GetBody() *GetOpsItemResponseBody {
	return s.Body
}

func (s *GetOpsItemResponse) SetHeaders(v map[string]*string) *GetOpsItemResponse {
	s.Headers = v
	return s
}

func (s *GetOpsItemResponse) SetStatusCode(v int32) *GetOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpsItemResponse) SetBody(v *GetOpsItemResponseBody) *GetOpsItemResponse {
	s.Body = v
	return s
}

func (s *GetOpsItemResponse) Validate() error {
	return dara.Validate(s)
}
