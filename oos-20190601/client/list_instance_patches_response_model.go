// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancePatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancePatchesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancePatchesResponseBody) *ListInstancePatchesResponse
	GetBody() *ListInstancePatchesResponseBody
}

type ListInstancePatchesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancePatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancePatchesResponse) GetBody() *ListInstancePatchesResponseBody {
	return s.Body
}

func (s *ListInstancePatchesResponse) SetHeaders(v map[string]*string) *ListInstancePatchesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchesResponse) SetStatusCode(v int32) *ListInstancePatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePatchesResponse) SetBody(v *ListInstancePatchesResponseBody) *ListInstancePatchesResponse {
	s.Body = v
	return s
}

func (s *ListInstancePatchesResponse) Validate() error {
	return dara.Validate(s)
}
