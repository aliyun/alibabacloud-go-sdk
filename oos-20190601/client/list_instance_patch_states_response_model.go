// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancePatchStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancePatchStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancePatchStatesResponseBody) *ListInstancePatchStatesResponse
	GetBody() *ListInstancePatchStatesResponseBody
}

type ListInstancePatchStatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePatchStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePatchStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchStatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancePatchStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancePatchStatesResponse) GetBody() *ListInstancePatchStatesResponseBody {
	return s.Body
}

func (s *ListInstancePatchStatesResponse) SetHeaders(v map[string]*string) *ListInstancePatchStatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchStatesResponse) SetStatusCode(v int32) *ListInstancePatchStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePatchStatesResponse) SetBody(v *ListInstancePatchStatesResponseBody) *ListInstancePatchStatesResponse {
	s.Body = v
	return s
}

func (s *ListInstancePatchStatesResponse) Validate() error {
	return dara.Validate(s)
}
