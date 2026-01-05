// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DisableDeletionProtectionResponseBody) *DisableDeletionProtectionResponse
	GetBody() *DisableDeletionProtectionResponseBody
}

type DisableDeletionProtectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDeletionProtectionResponse) GetBody() *DisableDeletionProtectionResponseBody {
	return s.Body
}

func (s *DisableDeletionProtectionResponse) SetHeaders(v map[string]*string) *DisableDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *DisableDeletionProtectionResponse) SetStatusCode(v int32) *DisableDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDeletionProtectionResponse) SetBody(v *DisableDeletionProtectionResponseBody) *DisableDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *DisableDeletionProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
