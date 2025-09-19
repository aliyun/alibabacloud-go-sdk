// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAssetSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAssetSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAssetSummaryResponseBody) *GetCloudAssetSummaryResponse
	GetBody() *GetCloudAssetSummaryResponseBody
}

type GetCloudAssetSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAssetSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAssetSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAssetSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAssetSummaryResponse) GetBody() *GetCloudAssetSummaryResponseBody {
	return s.Body
}

func (s *GetCloudAssetSummaryResponse) SetHeaders(v map[string]*string) *GetCloudAssetSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAssetSummaryResponse) SetStatusCode(v int32) *GetCloudAssetSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAssetSummaryResponse) SetBody(v *GetCloudAssetSummaryResponseBody) *GetCloudAssetSummaryResponse {
	s.Body = v
	return s
}

func (s *GetCloudAssetSummaryResponse) Validate() error {
	return dara.Validate(s)
}
