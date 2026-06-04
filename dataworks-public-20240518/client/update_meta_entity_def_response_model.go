// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityDefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaEntityDefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaEntityDefResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaEntityDefResponseBody) *UpdateMetaEntityDefResponse
	GetBody() *UpdateMetaEntityDefResponseBody
}

type UpdateMetaEntityDefResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaEntityDefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaEntityDefResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityDefResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityDefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaEntityDefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaEntityDefResponse) GetBody() *UpdateMetaEntityDefResponseBody {
	return s.Body
}

func (s *UpdateMetaEntityDefResponse) SetHeaders(v map[string]*string) *UpdateMetaEntityDefResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaEntityDefResponse) SetStatusCode(v int32) *UpdateMetaEntityDefResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaEntityDefResponse) SetBody(v *UpdateMetaEntityDefResponseBody) *UpdateMetaEntityDefResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaEntityDefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
