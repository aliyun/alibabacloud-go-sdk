// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomFieldResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomFieldResponseBody) *DeleteCustomFieldResponse
	GetBody() *DeleteCustomFieldResponseBody
}

type DeleteCustomFieldResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomFieldResponse) GetBody() *DeleteCustomFieldResponseBody {
	return s.Body
}

func (s *DeleteCustomFieldResponse) SetHeaders(v map[string]*string) *DeleteCustomFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomFieldResponse) SetStatusCode(v int32) *DeleteCustomFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomFieldResponse) SetBody(v *DeleteCustomFieldResponseBody) *DeleteCustomFieldResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
