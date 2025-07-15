// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStudioLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStudioLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStudioLayoutResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStudioLayoutResponseBody) *DeleteStudioLayoutResponse
	GetBody() *DeleteStudioLayoutResponseBody
}

type DeleteStudioLayoutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStudioLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStudioLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStudioLayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteStudioLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStudioLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStudioLayoutResponse) GetBody() *DeleteStudioLayoutResponseBody {
	return s.Body
}

func (s *DeleteStudioLayoutResponse) SetHeaders(v map[string]*string) *DeleteStudioLayoutResponse {
	s.Headers = v
	return s
}

func (s *DeleteStudioLayoutResponse) SetStatusCode(v int32) *DeleteStudioLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStudioLayoutResponse) SetBody(v *DeleteStudioLayoutResponseBody) *DeleteStudioLayoutResponse {
	s.Body = v
	return s
}

func (s *DeleteStudioLayoutResponse) Validate() error {
	return dara.Validate(s)
}
