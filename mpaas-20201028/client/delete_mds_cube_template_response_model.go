// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsCubeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMdsCubeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMdsCubeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMdsCubeTemplateResponseBody) *DeleteMdsCubeTemplateResponse
	GetBody() *DeleteMdsCubeTemplateResponseBody
}

type DeleteMdsCubeTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMdsCubeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMdsCubeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsCubeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteMdsCubeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMdsCubeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMdsCubeTemplateResponse) GetBody() *DeleteMdsCubeTemplateResponseBody {
	return s.Body
}

func (s *DeleteMdsCubeTemplateResponse) SetHeaders(v map[string]*string) *DeleteMdsCubeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteMdsCubeTemplateResponse) SetStatusCode(v int32) *DeleteMdsCubeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponse) SetBody(v *DeleteMdsCubeTemplateResponseBody) *DeleteMdsCubeTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteMdsCubeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
