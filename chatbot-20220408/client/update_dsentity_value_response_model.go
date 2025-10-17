// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDSEntityValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDSEntityValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDSEntityValueResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDSEntityValueResponseBody) *UpdateDSEntityValueResponse
	GetBody() *UpdateDSEntityValueResponseBody
}

type UpdateDSEntityValueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDSEntityValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDSEntityValueResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDSEntityValueResponse) GoString() string {
	return s.String()
}

func (s *UpdateDSEntityValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDSEntityValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDSEntityValueResponse) GetBody() *UpdateDSEntityValueResponseBody {
	return s.Body
}

func (s *UpdateDSEntityValueResponse) SetHeaders(v map[string]*string) *UpdateDSEntityValueResponse {
	s.Headers = v
	return s
}

func (s *UpdateDSEntityValueResponse) SetStatusCode(v int32) *UpdateDSEntityValueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDSEntityValueResponse) SetBody(v *UpdateDSEntityValueResponseBody) *UpdateDSEntityValueResponse {
	s.Body = v
	return s
}

func (s *UpdateDSEntityValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
