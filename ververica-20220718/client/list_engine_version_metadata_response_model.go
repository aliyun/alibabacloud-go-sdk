// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineVersionMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEngineVersionMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEngineVersionMetadataResponse
	GetStatusCode() *int32
	SetBody(v *ListEngineVersionMetadataResponseBody) *ListEngineVersionMetadataResponse
	GetBody() *ListEngineVersionMetadataResponseBody
}

type ListEngineVersionMetadataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEngineVersionMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEngineVersionMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEngineVersionMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEngineVersionMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEngineVersionMetadataResponse) GetBody() *ListEngineVersionMetadataResponseBody {
	return s.Body
}

func (s *ListEngineVersionMetadataResponse) SetHeaders(v map[string]*string) *ListEngineVersionMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListEngineVersionMetadataResponse) SetStatusCode(v int32) *ListEngineVersionMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponse) SetBody(v *ListEngineVersionMetadataResponseBody) *ListEngineVersionMetadataResponse {
	s.Body = v
	return s
}

func (s *ListEngineVersionMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
