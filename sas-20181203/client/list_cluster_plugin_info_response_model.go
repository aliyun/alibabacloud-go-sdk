// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterPluginInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterPluginInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterPluginInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterPluginInfoResponseBody) *ListClusterPluginInfoResponse
	GetBody() *ListClusterPluginInfoResponseBody
}

type ListClusterPluginInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterPluginInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterPluginInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterPluginInfoResponse) GoString() string {
	return s.String()
}

func (s *ListClusterPluginInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterPluginInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterPluginInfoResponse) GetBody() *ListClusterPluginInfoResponseBody {
	return s.Body
}

func (s *ListClusterPluginInfoResponse) SetHeaders(v map[string]*string) *ListClusterPluginInfoResponse {
	s.Headers = v
	return s
}

func (s *ListClusterPluginInfoResponse) SetStatusCode(v int32) *ListClusterPluginInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterPluginInfoResponse) SetBody(v *ListClusterPluginInfoResponseBody) *ListClusterPluginInfoResponse {
	s.Body = v
	return s
}

func (s *ListClusterPluginInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
