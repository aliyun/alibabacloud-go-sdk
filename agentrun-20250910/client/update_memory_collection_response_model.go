// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryCollectionResponse
	GetStatusCode() *int32
	SetBody(v *MemoryCollectionResult) *UpdateMemoryCollectionResponse
	GetBody() *MemoryCollectionResult
}

type UpdateMemoryCollectionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemoryCollectionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryCollectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryCollectionResponse) GetBody() *MemoryCollectionResult {
	return s.Body
}

func (s *UpdateMemoryCollectionResponse) SetHeaders(v map[string]*string) *UpdateMemoryCollectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryCollectionResponse) SetStatusCode(v int32) *UpdateMemoryCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryCollectionResponse) SetBody(v *MemoryCollectionResult) *UpdateMemoryCollectionResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
