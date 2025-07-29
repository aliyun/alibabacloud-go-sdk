// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAddonInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterAddonInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterAddonInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterAddonInstanceResponseBody) *GetClusterAddonInstanceResponse
	GetBody() *GetClusterAddonInstanceResponseBody
}

type GetClusterAddonInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterAddonInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterAddonInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAddonInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetClusterAddonInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterAddonInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterAddonInstanceResponse) GetBody() *GetClusterAddonInstanceResponseBody {
	return s.Body
}

func (s *GetClusterAddonInstanceResponse) SetHeaders(v map[string]*string) *GetClusterAddonInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetClusterAddonInstanceResponse) SetStatusCode(v int32) *GetClusterAddonInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterAddonInstanceResponse) SetBody(v *GetClusterAddonInstanceResponseBody) *GetClusterAddonInstanceResponse {
	s.Body = v
	return s
}

func (s *GetClusterAddonInstanceResponse) Validate() error {
	return dara.Validate(s)
}
