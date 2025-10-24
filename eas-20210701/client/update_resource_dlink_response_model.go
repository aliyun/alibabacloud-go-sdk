// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceDLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceDLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceDLinkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceDLinkResponseBody) *UpdateResourceDLinkResponse
	GetBody() *UpdateResourceDLinkResponseBody
}

type UpdateResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceDLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceDLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceDLinkResponse) GetBody() *UpdateResourceDLinkResponseBody {
	return s.Body
}

func (s *UpdateResourceDLinkResponse) SetHeaders(v map[string]*string) *UpdateResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceDLinkResponse) SetStatusCode(v int32) *UpdateResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceDLinkResponse) SetBody(v *UpdateResourceDLinkResponseBody) *UpdateResourceDLinkResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceDLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
