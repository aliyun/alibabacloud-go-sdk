// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetRefreshTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetRefreshTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetRefreshTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetRefreshTaskConfigResponseBody) *ListAssetRefreshTaskConfigResponse
	GetBody() *ListAssetRefreshTaskConfigResponseBody
}

type ListAssetRefreshTaskConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetRefreshTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetRefreshTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetRefreshTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *ListAssetRefreshTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetRefreshTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetRefreshTaskConfigResponse) GetBody() *ListAssetRefreshTaskConfigResponseBody {
	return s.Body
}

func (s *ListAssetRefreshTaskConfigResponse) SetHeaders(v map[string]*string) *ListAssetRefreshTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *ListAssetRefreshTaskConfigResponse) SetStatusCode(v int32) *ListAssetRefreshTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponse) SetBody(v *ListAssetRefreshTaskConfigResponseBody) *ListAssetRefreshTaskConfigResponse {
	s.Body = v
	return s
}

func (s *ListAssetRefreshTaskConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
