// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAssetRefreshTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAssetRefreshTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAssetRefreshTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAssetRefreshTaskConfigResponseBody) *ChangeAssetRefreshTaskConfigResponse
	GetBody() *ChangeAssetRefreshTaskConfigResponseBody
}

type ChangeAssetRefreshTaskConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAssetRefreshTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAssetRefreshTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAssetRefreshTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *ChangeAssetRefreshTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAssetRefreshTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAssetRefreshTaskConfigResponse) GetBody() *ChangeAssetRefreshTaskConfigResponseBody {
	return s.Body
}

func (s *ChangeAssetRefreshTaskConfigResponse) SetHeaders(v map[string]*string) *ChangeAssetRefreshTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponse) SetStatusCode(v int32) *ChangeAssetRefreshTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponse) SetBody(v *ChangeAssetRefreshTaskConfigResponseBody) *ChangeAssetRefreshTaskConfigResponse {
	s.Body = v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
