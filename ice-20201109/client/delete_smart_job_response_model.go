// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmartJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmartJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmartJobResponseBody) *DeleteSmartJobResponse
	GetBody() *DeleteSmartJobResponseBody
}

type DeleteSmartJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmartJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmartJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmartJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmartJobResponse) GetBody() *DeleteSmartJobResponseBody {
	return s.Body
}

func (s *DeleteSmartJobResponse) SetHeaders(v map[string]*string) *DeleteSmartJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmartJobResponse) SetStatusCode(v int32) *DeleteSmartJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmartJobResponse) SetBody(v *DeleteSmartJobResponseBody) *DeleteSmartJobResponse {
	s.Body = v
	return s
}

func (s *DeleteSmartJobResponse) Validate() error {
	return dara.Validate(s)
}
