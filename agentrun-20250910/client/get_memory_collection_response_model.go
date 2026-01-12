// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryCollectionResponse
	GetStatusCode() *int32
	SetBody(v *MemoryCollectionResult) *GetMemoryCollectionResponse
	GetBody() *MemoryCollectionResult
}

type GetMemoryCollectionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemoryCollectionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryCollectionResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryCollectionResponse) GetBody() *MemoryCollectionResult {
	return s.Body
}

func (s *GetMemoryCollectionResponse) SetHeaders(v map[string]*string) *GetMemoryCollectionResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryCollectionResponse) SetStatusCode(v int32) *GetMemoryCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryCollectionResponse) SetBody(v *MemoryCollectionResult) *GetMemoryCollectionResponse {
	s.Body = v
	return s
}

func (s *GetMemoryCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
