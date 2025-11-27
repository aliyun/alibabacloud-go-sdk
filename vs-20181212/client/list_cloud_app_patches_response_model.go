// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppPatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAppPatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAppPatchesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAppPatchesResponseBody) *ListCloudAppPatchesResponse
	GetBody() *ListCloudAppPatchesResponseBody
}

type ListCloudAppPatchesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAppPatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAppPatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppPatchesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAppPatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAppPatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAppPatchesResponse) GetBody() *ListCloudAppPatchesResponseBody {
	return s.Body
}

func (s *ListCloudAppPatchesResponse) SetHeaders(v map[string]*string) *ListCloudAppPatchesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAppPatchesResponse) SetStatusCode(v int32) *ListCloudAppPatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAppPatchesResponse) SetBody(v *ListCloudAppPatchesResponseBody) *ListCloudAppPatchesResponse {
	s.Body = v
	return s
}

func (s *ListCloudAppPatchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
