// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVscMountPointAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindVscMountPointAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindVscMountPointAliasResponse
	GetStatusCode() *int32
	SetBody(v *BindVscMountPointAliasResponseBody) *BindVscMountPointAliasResponse
	GetBody() *BindVscMountPointAliasResponseBody
}

type BindVscMountPointAliasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindVscMountPointAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindVscMountPointAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s BindVscMountPointAliasResponse) GoString() string {
	return s.String()
}

func (s *BindVscMountPointAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindVscMountPointAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindVscMountPointAliasResponse) GetBody() *BindVscMountPointAliasResponseBody {
	return s.Body
}

func (s *BindVscMountPointAliasResponse) SetHeaders(v map[string]*string) *BindVscMountPointAliasResponse {
	s.Headers = v
	return s
}

func (s *BindVscMountPointAliasResponse) SetStatusCode(v int32) *BindVscMountPointAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *BindVscMountPointAliasResponse) SetBody(v *BindVscMountPointAliasResponseBody) *BindVscMountPointAliasResponse {
	s.Body = v
	return s
}

func (s *BindVscMountPointAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
