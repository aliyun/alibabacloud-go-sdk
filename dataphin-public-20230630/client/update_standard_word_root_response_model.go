// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardWordRootResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardWordRootResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardWordRootResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardWordRootResponseBody) *UpdateStandardWordRootResponse
	GetBody() *UpdateStandardWordRootResponseBody
}

type UpdateStandardWordRootResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardWordRootResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardWordRootResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardWordRootResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardWordRootResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardWordRootResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardWordRootResponse) GetBody() *UpdateStandardWordRootResponseBody {
	return s.Body
}

func (s *UpdateStandardWordRootResponse) SetHeaders(v map[string]*string) *UpdateStandardWordRootResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardWordRootResponse) SetStatusCode(v int32) *UpdateStandardWordRootResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardWordRootResponse) SetBody(v *UpdateStandardWordRootResponseBody) *UpdateStandardWordRootResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardWordRootResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
