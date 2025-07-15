// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatchBaselinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPatchBaselinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPatchBaselinesResponse
	GetStatusCode() *int32
	SetBody(v *ListPatchBaselinesResponseBody) *ListPatchBaselinesResponse
	GetBody() *ListPatchBaselinesResponseBody
}

type ListPatchBaselinesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPatchBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPatchBaselinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPatchBaselinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPatchBaselinesResponse) GetBody() *ListPatchBaselinesResponseBody {
	return s.Body
}

func (s *ListPatchBaselinesResponse) SetHeaders(v map[string]*string) *ListPatchBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListPatchBaselinesResponse) SetStatusCode(v int32) *ListPatchBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPatchBaselinesResponse) SetBody(v *ListPatchBaselinesResponseBody) *ListPatchBaselinesResponse {
	s.Body = v
	return s
}

func (s *ListPatchBaselinesResponse) Validate() error {
	return dara.Validate(s)
}
