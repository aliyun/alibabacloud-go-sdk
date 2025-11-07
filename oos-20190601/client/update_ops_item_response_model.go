// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpsItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOpsItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOpsItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOpsItemResponseBody) *UpdateOpsItemResponse
	GetBody() *UpdateOpsItemResponseBody
}

type UpdateOpsItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOpsItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpsItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOpsItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOpsItemResponse) GetBody() *UpdateOpsItemResponseBody {
	return s.Body
}

func (s *UpdateOpsItemResponse) SetHeaders(v map[string]*string) *UpdateOpsItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateOpsItemResponse) SetStatusCode(v int32) *UpdateOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOpsItemResponse) SetBody(v *UpdateOpsItemResponseBody) *UpdateOpsItemResponse {
	s.Body = v
	return s
}

func (s *UpdateOpsItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
