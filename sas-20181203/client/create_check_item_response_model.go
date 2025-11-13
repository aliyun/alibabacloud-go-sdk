// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCheckItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCheckItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateCheckItemResponseBody) *CreateCheckItemResponse
	GetBody() *CreateCheckItemResponseBody
}

type CreateCheckItemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCheckItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCheckItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckItemResponse) GoString() string {
	return s.String()
}

func (s *CreateCheckItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCheckItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCheckItemResponse) GetBody() *CreateCheckItemResponseBody {
	return s.Body
}

func (s *CreateCheckItemResponse) SetHeaders(v map[string]*string) *CreateCheckItemResponse {
	s.Headers = v
	return s
}

func (s *CreateCheckItemResponse) SetStatusCode(v int32) *CreateCheckItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCheckItemResponse) SetBody(v *CreateCheckItemResponseBody) *CreateCheckItemResponse {
	s.Body = v
	return s
}

func (s *CreateCheckItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
