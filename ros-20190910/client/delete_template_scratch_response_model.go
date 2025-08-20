// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateScratchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateScratchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateScratchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateScratchResponseBody) *DeleteTemplateScratchResponse
	GetBody() *DeleteTemplateScratchResponseBody
}

type DeleteTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateScratchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateScratchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateScratchResponse) GetBody() *DeleteTemplateScratchResponseBody {
	return s.Body
}

func (s *DeleteTemplateScratchResponse) SetHeaders(v map[string]*string) *DeleteTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateScratchResponse) SetStatusCode(v int32) *DeleteTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateScratchResponse) SetBody(v *DeleteTemplateScratchResponseBody) *DeleteTemplateScratchResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateScratchResponse) Validate() error {
	return dara.Validate(s)
}
