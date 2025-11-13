// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCheckItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCheckItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCheckItemResponseBody) *UpdateCheckItemResponse
	GetBody() *UpdateCheckItemResponseBody
}

type UpdateCheckItemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCheckItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCheckItemResponse) GetBody() *UpdateCheckItemResponseBody {
	return s.Body
}

func (s *UpdateCheckItemResponse) SetHeaders(v map[string]*string) *UpdateCheckItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckItemResponse) SetStatusCode(v int32) *UpdateCheckItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckItemResponse) SetBody(v *UpdateCheckItemResponseBody) *UpdateCheckItemResponse {
	s.Body = v
	return s
}

func (s *UpdateCheckItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
