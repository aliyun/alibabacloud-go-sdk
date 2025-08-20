// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateScratchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateScratchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateScratchesResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateScratchesResponseBody) *ListTemplateScratchesResponse
	GetBody() *ListTemplateScratchesResponseBody
}

type ListTemplateScratchesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateScratchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateScratchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateScratchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateScratchesResponse) GetBody() *ListTemplateScratchesResponseBody {
	return s.Body
}

func (s *ListTemplateScratchesResponse) SetHeaders(v map[string]*string) *ListTemplateScratchesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateScratchesResponse) SetStatusCode(v int32) *ListTemplateScratchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateScratchesResponse) SetBody(v *ListTemplateScratchesResponseBody) *ListTemplateScratchesResponse {
	s.Body = v
	return s
}

func (s *ListTemplateScratchesResponse) Validate() error {
	return dara.Validate(s)
}
