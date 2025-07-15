// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocsResponseBody) *DeleteDocsResponse
	GetBody() *DeleteDocsResponseBody
}

type DeleteDocsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocsResponse) GetBody() *DeleteDocsResponseBody {
	return s.Body
}

func (s *DeleteDocsResponse) SetHeaders(v map[string]*string) *DeleteDocsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocsResponse) SetStatusCode(v int32) *DeleteDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocsResponse) SetBody(v *DeleteDocsResponseBody) *DeleteDocsResponse {
	s.Body = v
	return s
}

func (s *DeleteDocsResponse) Validate() error {
	return dara.Validate(s)
}
