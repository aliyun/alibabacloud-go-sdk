// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationVerifiedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappMigrationVerifiedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappMigrationVerifiedResponse
	GetStatusCode() *int32
	SetBody(v *ChatappMigrationVerifiedResponseBody) *ChatappMigrationVerifiedResponse
	GetBody() *ChatappMigrationVerifiedResponseBody
}

type ChatappMigrationVerifiedResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappMigrationVerifiedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappMigrationVerifiedResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationVerifiedResponse) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappMigrationVerifiedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappMigrationVerifiedResponse) GetBody() *ChatappMigrationVerifiedResponseBody {
	return s.Body
}

func (s *ChatappMigrationVerifiedResponse) SetHeaders(v map[string]*string) *ChatappMigrationVerifiedResponse {
	s.Headers = v
	return s
}

func (s *ChatappMigrationVerifiedResponse) SetStatusCode(v int32) *ChatappMigrationVerifiedResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappMigrationVerifiedResponse) SetBody(v *ChatappMigrationVerifiedResponseBody) *ChatappMigrationVerifiedResponse {
	s.Body = v
	return s
}

func (s *ChatappMigrationVerifiedResponse) Validate() error {
	return dara.Validate(s)
}
