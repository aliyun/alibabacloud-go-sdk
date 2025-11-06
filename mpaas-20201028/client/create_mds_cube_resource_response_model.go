// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMdsCubeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMdsCubeResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMdsCubeResourceResponseBody) *CreateMdsCubeResourceResponse
	GetBody() *CreateMdsCubeResourceResponseBody
}

type CreateMdsCubeResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMdsCubeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMdsCubeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMdsCubeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMdsCubeResourceResponse) GetBody() *CreateMdsCubeResourceResponseBody {
	return s.Body
}

func (s *CreateMdsCubeResourceResponse) SetHeaders(v map[string]*string) *CreateMdsCubeResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateMdsCubeResourceResponse) SetStatusCode(v int32) *CreateMdsCubeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMdsCubeResourceResponse) SetBody(v *CreateMdsCubeResourceResponseBody) *CreateMdsCubeResourceResponse {
	s.Body = v
	return s
}

func (s *CreateMdsCubeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
