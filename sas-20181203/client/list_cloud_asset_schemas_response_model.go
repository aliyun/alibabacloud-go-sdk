// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAssetSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAssetSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAssetSchemasResponseBody) *ListCloudAssetSchemasResponse
	GetBody() *ListCloudAssetSchemasResponseBody
}

type ListCloudAssetSchemasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAssetSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAssetSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAssetSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAssetSchemasResponse) GetBody() *ListCloudAssetSchemasResponseBody {
	return s.Body
}

func (s *ListCloudAssetSchemasResponse) SetHeaders(v map[string]*string) *ListCloudAssetSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAssetSchemasResponse) SetStatusCode(v int32) *ListCloudAssetSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAssetSchemasResponse) SetBody(v *ListCloudAssetSchemasResponseBody) *ListCloudAssetSchemasResponse {
	s.Body = v
	return s
}

func (s *ListCloudAssetSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
