// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddonTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddonTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListAddonTemplatesResponseBody) *ListAddonTemplatesResponse
	GetBody() *ListAddonTemplatesResponseBody
}

type ListAddonTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddonTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddonTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddonTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddonTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddonTemplatesResponse) GetBody() *ListAddonTemplatesResponseBody {
	return s.Body
}

func (s *ListAddonTemplatesResponse) SetHeaders(v map[string]*string) *ListAddonTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAddonTemplatesResponse) SetStatusCode(v int32) *ListAddonTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddonTemplatesResponse) SetBody(v *ListAddonTemplatesResponseBody) *ListAddonTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListAddonTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
