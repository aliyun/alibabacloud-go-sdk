// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaRestoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHanaRestoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHanaRestoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateHanaRestoreResponseBody) *CreateHanaRestoreResponse
	GetBody() *CreateHanaRestoreResponseBody
}

type CreateHanaRestoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHanaRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHanaRestoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaRestoreResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHanaRestoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHanaRestoreResponse) GetBody() *CreateHanaRestoreResponseBody {
	return s.Body
}

func (s *CreateHanaRestoreResponse) SetHeaders(v map[string]*string) *CreateHanaRestoreResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaRestoreResponse) SetStatusCode(v int32) *CreateHanaRestoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaRestoreResponse) SetBody(v *CreateHanaRestoreResponseBody) *CreateHanaRestoreResponse {
	s.Body = v
	return s
}

func (s *CreateHanaRestoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
