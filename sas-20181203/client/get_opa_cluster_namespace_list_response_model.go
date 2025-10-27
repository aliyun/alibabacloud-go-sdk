// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterNamespaceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaClusterNamespaceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaClusterNamespaceListResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaClusterNamespaceListResponseBody) *GetOpaClusterNamespaceListResponse
	GetBody() *GetOpaClusterNamespaceListResponseBody
}

type GetOpaClusterNamespaceListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaClusterNamespaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaClusterNamespaceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterNamespaceListResponse) GoString() string {
	return s.String()
}

func (s *GetOpaClusterNamespaceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaClusterNamespaceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaClusterNamespaceListResponse) GetBody() *GetOpaClusterNamespaceListResponseBody {
	return s.Body
}

func (s *GetOpaClusterNamespaceListResponse) SetHeaders(v map[string]*string) *GetOpaClusterNamespaceListResponse {
	s.Headers = v
	return s
}

func (s *GetOpaClusterNamespaceListResponse) SetStatusCode(v int32) *GetOpaClusterNamespaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaClusterNamespaceListResponse) SetBody(v *GetOpaClusterNamespaceListResponseBody) *GetOpaClusterNamespaceListResponse {
	s.Body = v
	return s
}

func (s *GetOpaClusterNamespaceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
