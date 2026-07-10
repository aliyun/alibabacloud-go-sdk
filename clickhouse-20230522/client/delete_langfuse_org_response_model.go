// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLangfuseOrgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLangfuseOrgResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLangfuseOrgResponseBody) *DeleteLangfuseOrgResponse
	GetBody() *DeleteLangfuseOrgResponseBody
}

type DeleteLangfuseOrgResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLangfuseOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLangfuseOrgResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgResponse) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLangfuseOrgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLangfuseOrgResponse) GetBody() *DeleteLangfuseOrgResponseBody {
	return s.Body
}

func (s *DeleteLangfuseOrgResponse) SetHeaders(v map[string]*string) *DeleteLangfuseOrgResponse {
	s.Headers = v
	return s
}

func (s *DeleteLangfuseOrgResponse) SetStatusCode(v int32) *DeleteLangfuseOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLangfuseOrgResponse) SetBody(v *DeleteLangfuseOrgResponseBody) *DeleteLangfuseOrgResponse {
	s.Body = v
	return s
}

func (s *DeleteLangfuseOrgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
