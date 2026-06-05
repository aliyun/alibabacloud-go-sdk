// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHiveResponse
	GetStatusCode() *int32
	SetBody(v *CreateHiveResponseBody) *CreateHiveResponse
	GetBody() *CreateHiveResponseBody
}

type CreateHiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHiveResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHiveResponse) GoString() string {
	return s.String()
}

func (s *CreateHiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHiveResponse) GetBody() *CreateHiveResponseBody {
	return s.Body
}

func (s *CreateHiveResponse) SetHeaders(v map[string]*string) *CreateHiveResponse {
	s.Headers = v
	return s
}

func (s *CreateHiveResponse) SetStatusCode(v int32) *CreateHiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHiveResponse) SetBody(v *CreateHiveResponseBody) *CreateHiveResponse {
	s.Body = v
	return s
}

func (s *CreateHiveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
