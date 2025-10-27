// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskLogDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSyncAssetTaskLogDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSyncAssetTaskLogDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSyncAssetTaskLogDetailResponseBody) *DescribeSyncAssetTaskLogDetailResponse
	GetBody() *DescribeSyncAssetTaskLogDetailResponseBody
}

type DescribeSyncAssetTaskLogDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncAssetTaskLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncAssetTaskLogDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskLogDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskLogDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSyncAssetTaskLogDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSyncAssetTaskLogDetailResponse) GetBody() *DescribeSyncAssetTaskLogDetailResponseBody {
	return s.Body
}

func (s *DescribeSyncAssetTaskLogDetailResponse) SetHeaders(v map[string]*string) *DescribeSyncAssetTaskLogDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponse) SetStatusCode(v int32) *DescribeSyncAssetTaskLogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponse) SetBody(v *DescribeSyncAssetTaskLogDetailResponseBody) *DescribeSyncAssetTaskLogDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
