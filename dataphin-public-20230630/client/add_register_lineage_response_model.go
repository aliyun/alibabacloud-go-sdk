// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRegisterLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRegisterLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRegisterLineageResponse
	GetStatusCode() *int32
	SetBody(v *AddRegisterLineageResponseBody) *AddRegisterLineageResponse
	GetBody() *AddRegisterLineageResponseBody
}

type AddRegisterLineageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRegisterLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRegisterLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageResponse) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRegisterLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRegisterLineageResponse) GetBody() *AddRegisterLineageResponseBody {
	return s.Body
}

func (s *AddRegisterLineageResponse) SetHeaders(v map[string]*string) *AddRegisterLineageResponse {
	s.Headers = v
	return s
}

func (s *AddRegisterLineageResponse) SetStatusCode(v int32) *AddRegisterLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRegisterLineageResponse) SetBody(v *AddRegisterLineageResponseBody) *AddRegisterLineageResponse {
	s.Body = v
	return s
}

func (s *AddRegisterLineageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
