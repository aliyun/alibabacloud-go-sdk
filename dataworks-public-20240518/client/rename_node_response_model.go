// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameNodeResponse
	GetStatusCode() *int32
	SetBody(v *RenameNodeResponseBody) *RenameNodeResponse
	GetBody() *RenameNodeResponseBody
}

type RenameNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameNodeResponse) GoString() string {
	return s.String()
}

func (s *RenameNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameNodeResponse) GetBody() *RenameNodeResponseBody {
	return s.Body
}

func (s *RenameNodeResponse) SetHeaders(v map[string]*string) *RenameNodeResponse {
	s.Headers = v
	return s
}

func (s *RenameNodeResponse) SetStatusCode(v int32) *RenameNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameNodeResponse) SetBody(v *RenameNodeResponseBody) *RenameNodeResponse {
	s.Body = v
	return s
}

func (s *RenameNodeResponse) Validate() error {
	return dara.Validate(s)
}
