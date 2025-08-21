// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSDGResponse
	GetStatusCode() *int32
	SetBody(v *SaveSDGResponseBody) *SaveSDGResponse
	GetBody() *SaveSDGResponseBody
}

type SaveSDGResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSDGResponse) GoString() string {
	return s.String()
}

func (s *SaveSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSDGResponse) GetBody() *SaveSDGResponseBody {
	return s.Body
}

func (s *SaveSDGResponse) SetHeaders(v map[string]*string) *SaveSDGResponse {
	s.Headers = v
	return s
}

func (s *SaveSDGResponse) SetStatusCode(v int32) *SaveSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSDGResponse) SetBody(v *SaveSDGResponseBody) *SaveSDGResponse {
	s.Body = v
	return s
}

func (s *SaveSDGResponse) Validate() error {
	return dara.Validate(s)
}
