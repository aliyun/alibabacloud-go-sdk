// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetWatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAssetWatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAssetWatchResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *CreateAssetWatchResponse
	GetBody() *CatalogCommonResult
}

type CreateAssetWatchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssetWatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetWatchResponse) GoString() string {
	return s.String()
}

func (s *CreateAssetWatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAssetWatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAssetWatchResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *CreateAssetWatchResponse) SetHeaders(v map[string]*string) *CreateAssetWatchResponse {
	s.Headers = v
	return s
}

func (s *CreateAssetWatchResponse) SetStatusCode(v int32) *CreateAssetWatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssetWatchResponse) SetBody(v *CatalogCommonResult) *CreateAssetWatchResponse {
	s.Body = v
	return s
}

func (s *CreateAssetWatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
