// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudVendorTrialConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudVendorTrialConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudVendorTrialConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudVendorTrialConfigResponseBody) *ModifyCloudVendorTrialConfigResponse
	GetBody() *ModifyCloudVendorTrialConfigResponseBody
}

type ModifyCloudVendorTrialConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudVendorTrialConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudVendorTrialConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorTrialConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorTrialConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudVendorTrialConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudVendorTrialConfigResponse) GetBody() *ModifyCloudVendorTrialConfigResponseBody {
	return s.Body
}

func (s *ModifyCloudVendorTrialConfigResponse) SetHeaders(v map[string]*string) *ModifyCloudVendorTrialConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudVendorTrialConfigResponse) SetStatusCode(v int32) *ModifyCloudVendorTrialConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudVendorTrialConfigResponse) SetBody(v *ModifyCloudVendorTrialConfigResponseBody) *ModifyCloudVendorTrialConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudVendorTrialConfigResponse) Validate() error {
	return dara.Validate(s)
}
