// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreTableResponse
	GetStatusCode() *int32
	SetBody(v *RestoreTableResponseBody) *RestoreTableResponse
	GetBody() *RestoreTableResponseBody
}

type RestoreTableResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreTableResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreTableResponse) GoString() string {
	return s.String()
}

func (s *RestoreTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreTableResponse) GetBody() *RestoreTableResponseBody {
	return s.Body
}

func (s *RestoreTableResponse) SetHeaders(v map[string]*string) *RestoreTableResponse {
	s.Headers = v
	return s
}

func (s *RestoreTableResponse) SetStatusCode(v int32) *RestoreTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreTableResponse) SetBody(v *RestoreTableResponseBody) *RestoreTableResponse {
	s.Body = v
	return s
}

func (s *RestoreTableResponse) Validate() error {
	return dara.Validate(s)
}
