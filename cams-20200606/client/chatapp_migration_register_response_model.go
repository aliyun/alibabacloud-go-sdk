// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationRegisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappMigrationRegisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappMigrationRegisterResponse
	GetStatusCode() *int32
	SetBody(v *ChatappMigrationRegisterResponseBody) *ChatappMigrationRegisterResponse
	GetBody() *ChatappMigrationRegisterResponseBody
}

type ChatappMigrationRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappMigrationRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappMigrationRegisterResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappMigrationRegisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappMigrationRegisterResponse) GetBody() *ChatappMigrationRegisterResponseBody {
	return s.Body
}

func (s *ChatappMigrationRegisterResponse) SetHeaders(v map[string]*string) *ChatappMigrationRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappMigrationRegisterResponse) SetStatusCode(v int32) *ChatappMigrationRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappMigrationRegisterResponse) SetBody(v *ChatappMigrationRegisterResponseBody) *ChatappMigrationRegisterResponse {
	s.Body = v
	return s
}

func (s *ChatappMigrationRegisterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
