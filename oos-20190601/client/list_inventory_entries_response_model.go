// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInventoryEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInventoryEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListInventoryEntriesResponseBody) *ListInventoryEntriesResponse
	GetBody() *ListInventoryEntriesResponseBody
}

type ListInventoryEntriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInventoryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInventoryEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInventoryEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInventoryEntriesResponse) GetBody() *ListInventoryEntriesResponseBody {
	return s.Body
}

func (s *ListInventoryEntriesResponse) SetHeaders(v map[string]*string) *ListInventoryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInventoryEntriesResponse) SetStatusCode(v int32) *ListInventoryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInventoryEntriesResponse) SetBody(v *ListInventoryEntriesResponseBody) *ListInventoryEntriesResponse {
	s.Body = v
	return s
}

func (s *ListInventoryEntriesResponse) Validate() error {
	return dara.Validate(s)
}
