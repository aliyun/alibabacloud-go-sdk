// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigDirResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAdvanceConfigDirResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAdvanceConfigDirResponse
	GetStatusCode() *int32
	SetBody(v *ListAdvanceConfigDirResponseBody) *ListAdvanceConfigDirResponse
	GetBody() *ListAdvanceConfigDirResponseBody
}

type ListAdvanceConfigDirResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdvanceConfigDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdvanceConfigDirResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigDirResponse) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAdvanceConfigDirResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAdvanceConfigDirResponse) GetBody() *ListAdvanceConfigDirResponseBody {
	return s.Body
}

func (s *ListAdvanceConfigDirResponse) SetHeaders(v map[string]*string) *ListAdvanceConfigDirResponse {
	s.Headers = v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetStatusCode(v int32) *ListAdvanceConfigDirResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdvanceConfigDirResponse) SetBody(v *ListAdvanceConfigDirResponseBody) *ListAdvanceConfigDirResponse {
	s.Body = v
	return s
}

func (s *ListAdvanceConfigDirResponse) Validate() error {
	return dara.Validate(s)
}
