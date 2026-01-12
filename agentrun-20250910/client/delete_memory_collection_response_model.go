// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoryCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoryCollectionResponse
	GetStatusCode() *int32
	SetBody(v *MemoryCollectionResult) *DeleteMemoryCollectionResponse
	GetBody() *MemoryCollectionResult
}

type DeleteMemoryCollectionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemoryCollectionResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoryCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoryCollectionResponse) GetBody() *MemoryCollectionResult {
	return s.Body
}

func (s *DeleteMemoryCollectionResponse) SetHeaders(v map[string]*string) *DeleteMemoryCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryCollectionResponse) SetStatusCode(v int32) *DeleteMemoryCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryCollectionResponse) SetBody(v *MemoryCollectionResult) *DeleteMemoryCollectionResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoryCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
