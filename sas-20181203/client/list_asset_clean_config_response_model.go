// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCleanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetCleanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetCleanConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetCleanConfigResponseBody) *ListAssetCleanConfigResponse
	GetBody() *ListAssetCleanConfigResponseBody
}

type ListAssetCleanConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetCleanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetCleanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCleanConfigResponse) GoString() string {
	return s.String()
}

func (s *ListAssetCleanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetCleanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetCleanConfigResponse) GetBody() *ListAssetCleanConfigResponseBody {
	return s.Body
}

func (s *ListAssetCleanConfigResponse) SetHeaders(v map[string]*string) *ListAssetCleanConfigResponse {
	s.Headers = v
	return s
}

func (s *ListAssetCleanConfigResponse) SetStatusCode(v int32) *ListAssetCleanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetCleanConfigResponse) SetBody(v *ListAssetCleanConfigResponseBody) *ListAssetCleanConfigResponse {
	s.Body = v
	return s
}

func (s *ListAssetCleanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
