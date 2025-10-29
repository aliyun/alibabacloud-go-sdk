// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseManagedIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseManagedIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseManagedIndexResponse
	GetStatusCode() *int32
	SetBody(v *CloseManagedIndexResponseBody) *CloseManagedIndexResponse
	GetBody() *CloseManagedIndexResponseBody
}

type CloseManagedIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseManagedIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseManagedIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseManagedIndexResponse) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseManagedIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseManagedIndexResponse) GetBody() *CloseManagedIndexResponseBody {
	return s.Body
}

func (s *CloseManagedIndexResponse) SetHeaders(v map[string]*string) *CloseManagedIndexResponse {
	s.Headers = v
	return s
}

func (s *CloseManagedIndexResponse) SetStatusCode(v int32) *CloseManagedIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseManagedIndexResponse) SetBody(v *CloseManagedIndexResponseBody) *CloseManagedIndexResponse {
	s.Body = v
	return s
}

func (s *CloseManagedIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
