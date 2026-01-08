// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappMigrationInitiateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatappMigrationInitiateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatappMigrationInitiateResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatappMigrationInitiateResponseBody) *CreateChatappMigrationInitiateResponse
	GetBody() *CreateChatappMigrationInitiateResponseBody
}

type CreateChatappMigrationInitiateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatappMigrationInitiateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatappMigrationInitiateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappMigrationInitiateResponse) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatappMigrationInitiateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatappMigrationInitiateResponse) GetBody() *CreateChatappMigrationInitiateResponseBody {
	return s.Body
}

func (s *CreateChatappMigrationInitiateResponse) SetHeaders(v map[string]*string) *CreateChatappMigrationInitiateResponse {
	s.Headers = v
	return s
}

func (s *CreateChatappMigrationInitiateResponse) SetStatusCode(v int32) *CreateChatappMigrationInitiateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatappMigrationInitiateResponse) SetBody(v *CreateChatappMigrationInitiateResponseBody) *CreateChatappMigrationInitiateResponse {
	s.Body = v
	return s
}

func (s *CreateChatappMigrationInitiateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
