// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerScanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContainerScanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContainerScanConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContainerScanConfigResponseBody) *ModifyContainerScanConfigResponse
	GetBody() *ModifyContainerScanConfigResponseBody
}

type ModifyContainerScanConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContainerScanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContainerScanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerScanConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerScanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContainerScanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContainerScanConfigResponse) GetBody() *ModifyContainerScanConfigResponseBody {
	return s.Body
}

func (s *ModifyContainerScanConfigResponse) SetHeaders(v map[string]*string) *ModifyContainerScanConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerScanConfigResponse) SetStatusCode(v int32) *ModifyContainerScanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContainerScanConfigResponse) SetBody(v *ModifyContainerScanConfigResponseBody) *ModifyContainerScanConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyContainerScanConfigResponse) Validate() error {
	return dara.Validate(s)
}
