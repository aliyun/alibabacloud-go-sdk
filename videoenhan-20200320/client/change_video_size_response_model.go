// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVideoSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeVideoSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeVideoSizeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeVideoSizeResponseBody) *ChangeVideoSizeResponse
	GetBody() *ChangeVideoSizeResponseBody
}

type ChangeVideoSizeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeVideoSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeVideoSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeVideoSizeResponse) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeVideoSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeVideoSizeResponse) GetBody() *ChangeVideoSizeResponseBody {
	return s.Body
}

func (s *ChangeVideoSizeResponse) SetHeaders(v map[string]*string) *ChangeVideoSizeResponse {
	s.Headers = v
	return s
}

func (s *ChangeVideoSizeResponse) SetStatusCode(v int32) *ChangeVideoSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeVideoSizeResponse) SetBody(v *ChangeVideoSizeResponseBody) *ChangeVideoSizeResponse {
	s.Body = v
	return s
}

func (s *ChangeVideoSizeResponse) Validate() error {
	return dara.Validate(s)
}
