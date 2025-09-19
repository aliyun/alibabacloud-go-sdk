// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorTrialConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCloudVendorTrialConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCloudVendorTrialConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddCloudVendorTrialConfigResponseBody) *AddCloudVendorTrialConfigResponse
	GetBody() *AddCloudVendorTrialConfigResponseBody
}

type AddCloudVendorTrialConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCloudVendorTrialConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCloudVendorTrialConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorTrialConfigResponse) GoString() string {
	return s.String()
}

func (s *AddCloudVendorTrialConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCloudVendorTrialConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCloudVendorTrialConfigResponse) GetBody() *AddCloudVendorTrialConfigResponseBody {
	return s.Body
}

func (s *AddCloudVendorTrialConfigResponse) SetHeaders(v map[string]*string) *AddCloudVendorTrialConfigResponse {
	s.Headers = v
	return s
}

func (s *AddCloudVendorTrialConfigResponse) SetStatusCode(v int32) *AddCloudVendorTrialConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCloudVendorTrialConfigResponse) SetBody(v *AddCloudVendorTrialConfigResponseBody) *AddCloudVendorTrialConfigResponse {
	s.Body = v
	return s
}

func (s *AddCloudVendorTrialConfigResponse) Validate() error {
	return dara.Validate(s)
}
