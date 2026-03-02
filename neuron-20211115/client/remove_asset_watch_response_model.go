// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAssetWatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAssetWatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAssetWatchResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *RemoveAssetWatchResponse
	GetBody() *CatalogCommonResult
}

type RemoveAssetWatchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAssetWatchResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAssetWatchResponse) GoString() string {
	return s.String()
}

func (s *RemoveAssetWatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAssetWatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAssetWatchResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *RemoveAssetWatchResponse) SetHeaders(v map[string]*string) *RemoveAssetWatchResponse {
	s.Headers = v
	return s
}

func (s *RemoveAssetWatchResponse) SetStatusCode(v int32) *RemoveAssetWatchResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAssetWatchResponse) SetBody(v *CatalogCommonResult) *RemoveAssetWatchResponse {
	s.Body = v
	return s
}

func (s *RemoveAssetWatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
