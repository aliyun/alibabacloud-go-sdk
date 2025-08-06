// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageRegionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageRegionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageRegionListResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageRegionListResponseBody) *GetStorageRegionListResponse
	GetBody() *GetStorageRegionListResponseBody
}

type GetStorageRegionListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageRegionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRegionListResponse) GoString() string {
	return s.String()
}

func (s *GetStorageRegionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageRegionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageRegionListResponse) GetBody() *GetStorageRegionListResponseBody {
	return s.Body
}

func (s *GetStorageRegionListResponse) SetHeaders(v map[string]*string) *GetStorageRegionListResponse {
	s.Headers = v
	return s
}

func (s *GetStorageRegionListResponse) SetStatusCode(v int32) *GetStorageRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageRegionListResponse) SetBody(v *GetStorageRegionListResponseBody) *GetStorageRegionListResponse {
	s.Body = v
	return s
}

func (s *GetStorageRegionListResponse) Validate() error {
	return dara.Validate(s)
}
