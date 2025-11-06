// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMdsCubeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMdsCubeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMdsCubeResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMdsCubeResourceResponseBody) *UpdateMdsCubeResourceResponse
	GetBody() *UpdateMdsCubeResourceResponseBody
}

type UpdateMdsCubeResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMdsCubeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMdsCubeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMdsCubeResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateMdsCubeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMdsCubeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMdsCubeResourceResponse) GetBody() *UpdateMdsCubeResourceResponseBody {
	return s.Body
}

func (s *UpdateMdsCubeResourceResponse) SetHeaders(v map[string]*string) *UpdateMdsCubeResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateMdsCubeResourceResponse) SetStatusCode(v int32) *UpdateMdsCubeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMdsCubeResourceResponse) SetBody(v *UpdateMdsCubeResourceResponseBody) *UpdateMdsCubeResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateMdsCubeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
