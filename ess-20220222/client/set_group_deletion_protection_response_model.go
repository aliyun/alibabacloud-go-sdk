// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetGroupDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetGroupDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetGroupDeletionProtectionResponseBody) *SetGroupDeletionProtectionResponse
	GetBody() *SetGroupDeletionProtectionResponseBody
}

type SetGroupDeletionProtectionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetGroupDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetGroupDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetGroupDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetGroupDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetGroupDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetGroupDeletionProtectionResponse) GetBody() *SetGroupDeletionProtectionResponseBody {
	return s.Body
}

func (s *SetGroupDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetGroupDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetGroupDeletionProtectionResponse) SetStatusCode(v int32) *SetGroupDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGroupDeletionProtectionResponse) SetBody(v *SetGroupDeletionProtectionResponseBody) *SetGroupDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *SetGroupDeletionProtectionResponse) Validate() error {
	return dara.Validate(s)
}
