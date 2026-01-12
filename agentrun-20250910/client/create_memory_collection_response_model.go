// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryCollectionResponse
	GetStatusCode() *int32
	SetBody(v *MemoryCollectionResult) *CreateMemoryCollectionResponse
	GetBody() *MemoryCollectionResult
}

type CreateMemoryCollectionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemoryCollectionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryCollectionResponse) GetBody() *MemoryCollectionResult {
	return s.Body
}

func (s *CreateMemoryCollectionResponse) SetHeaders(v map[string]*string) *CreateMemoryCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryCollectionResponse) SetStatusCode(v int32) *CreateMemoryCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryCollectionResponse) SetBody(v *MemoryCollectionResult) *CreateMemoryCollectionResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
