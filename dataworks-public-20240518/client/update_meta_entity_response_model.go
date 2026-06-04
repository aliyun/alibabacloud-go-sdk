// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMetaEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMetaEntityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMetaEntityResponseBody) *UpdateMetaEntityResponse
	GetBody() *UpdateMetaEntityResponseBody
}

type UpdateMetaEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMetaEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMetaEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMetaEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMetaEntityResponse) GetBody() *UpdateMetaEntityResponseBody {
	return s.Body
}

func (s *UpdateMetaEntityResponse) SetHeaders(v map[string]*string) *UpdateMetaEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateMetaEntityResponse) SetStatusCode(v int32) *UpdateMetaEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMetaEntityResponse) SetBody(v *UpdateMetaEntityResponseBody) *UpdateMetaEntityResponse {
	s.Body = v
	return s
}

func (s *UpdateMetaEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
