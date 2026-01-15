// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IncreaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IncreaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *IncreaseInstanceResponseBody) *IncreaseInstanceResponse
	GetBody() *IncreaseInstanceResponseBody
}

type IncreaseInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncreaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncreaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s IncreaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IncreaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IncreaseInstanceResponse) GetBody() *IncreaseInstanceResponseBody {
	return s.Body
}

func (s *IncreaseInstanceResponse) SetHeaders(v map[string]*string) *IncreaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *IncreaseInstanceResponse) SetStatusCode(v int32) *IncreaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseInstanceResponse) SetBody(v *IncreaseInstanceResponseBody) *IncreaseInstanceResponse {
	s.Body = v
	return s
}

func (s *IncreaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
