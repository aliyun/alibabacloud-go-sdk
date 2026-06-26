// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsByNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPartitionsByNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPartitionsByNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListPartitionsByNamesResponseBody) *ListPartitionsByNamesResponse
	GetBody() *ListPartitionsByNamesResponseBody
}

type ListPartitionsByNamesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPartitionsByNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPartitionsByNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsByNamesResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionsByNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPartitionsByNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPartitionsByNamesResponse) GetBody() *ListPartitionsByNamesResponseBody {
	return s.Body
}

func (s *ListPartitionsByNamesResponse) SetHeaders(v map[string]*string) *ListPartitionsByNamesResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionsByNamesResponse) SetStatusCode(v int32) *ListPartitionsByNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionsByNamesResponse) SetBody(v *ListPartitionsByNamesResponseBody) *ListPartitionsByNamesResponse {
	s.Body = v
	return s
}

func (s *ListPartitionsByNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
