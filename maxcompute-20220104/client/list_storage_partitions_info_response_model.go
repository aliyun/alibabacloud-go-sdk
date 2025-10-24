// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStoragePartitionsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStoragePartitionsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStoragePartitionsInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListStoragePartitionsInfoResponseBody) *ListStoragePartitionsInfoResponse
	GetBody() *ListStoragePartitionsInfoResponseBody
}

type ListStoragePartitionsInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStoragePartitionsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStoragePartitionsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponse) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStoragePartitionsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStoragePartitionsInfoResponse) GetBody() *ListStoragePartitionsInfoResponseBody {
	return s.Body
}

func (s *ListStoragePartitionsInfoResponse) SetHeaders(v map[string]*string) *ListStoragePartitionsInfoResponse {
	s.Headers = v
	return s
}

func (s *ListStoragePartitionsInfoResponse) SetStatusCode(v int32) *ListStoragePartitionsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponse) SetBody(v *ListStoragePartitionsInfoResponseBody) *ListStoragePartitionsInfoResponse {
	s.Body = v
	return s
}

func (s *ListStoragePartitionsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
