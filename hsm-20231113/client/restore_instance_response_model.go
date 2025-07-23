// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestoreInstanceResponseBody) *RestoreInstanceResponse
	GetBody() *RestoreInstanceResponseBody
}

type RestoreInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreInstanceResponse) GetBody() *RestoreInstanceResponseBody {
	return s.Body
}

func (s *RestoreInstanceResponse) SetHeaders(v map[string]*string) *RestoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestoreInstanceResponse) SetStatusCode(v int32) *RestoreInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreInstanceResponse) SetBody(v *RestoreInstanceResponseBody) *RestoreInstanceResponse {
	s.Body = v
	return s
}

func (s *RestoreInstanceResponse) Validate() error {
	return dara.Validate(s)
}
