// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetDeletionProtectionResponseBody) *SetDeletionProtectionResponse
	GetBody() *SetDeletionProtectionResponseBody
}

type SetDeletionProtectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDeletionProtectionResponse) GetBody() *SetDeletionProtectionResponseBody {
	return s.Body
}

func (s *SetDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetDeletionProtectionResponse) SetStatusCode(v int32) *SetDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeletionProtectionResponse) SetBody(v *SetDeletionProtectionResponseBody) *SetDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *SetDeletionProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
