// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DeletionProtectionResponseBody) *DeletionProtectionResponse
	GetBody() *DeletionProtectionResponseBody
}

type DeletionProtectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *DeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletionProtectionResponse) GetBody() *DeletionProtectionResponseBody {
	return s.Body
}

func (s *DeletionProtectionResponse) SetHeaders(v map[string]*string) *DeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *DeletionProtectionResponse) SetStatusCode(v int32) *DeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletionProtectionResponse) SetBody(v *DeletionProtectionResponseBody) *DeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *DeletionProtectionResponse) Validate() error {
	return dara.Validate(s)
}
