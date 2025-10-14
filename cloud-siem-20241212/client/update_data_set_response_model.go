// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSetResponseBody) *UpdateDataSetResponse
	GetBody() *UpdateDataSetResponseBody
}

type UpdateDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSetResponse) GetBody() *UpdateDataSetResponseBody {
	return s.Body
}

func (s *UpdateDataSetResponse) SetHeaders(v map[string]*string) *UpdateDataSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSetResponse) SetStatusCode(v int32) *UpdateDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSetResponse) SetBody(v *UpdateDataSetResponseBody) *UpdateDataSetResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
