// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameResourceResponse
	GetStatusCode() *int32
	SetBody(v *RenameResourceResponseBody) *RenameResourceResponse
	GetBody() *RenameResourceResponseBody
}

type RenameResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameResourceResponse) GoString() string {
	return s.String()
}

func (s *RenameResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameResourceResponse) GetBody() *RenameResourceResponseBody {
	return s.Body
}

func (s *RenameResourceResponse) SetHeaders(v map[string]*string) *RenameResourceResponse {
	s.Headers = v
	return s
}

func (s *RenameResourceResponse) SetStatusCode(v int32) *RenameResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameResourceResponse) SetBody(v *RenameResourceResponseBody) *RenameResourceResponse {
	s.Body = v
	return s
}

func (s *RenameResourceResponse) Validate() error {
	return dara.Validate(s)
}
