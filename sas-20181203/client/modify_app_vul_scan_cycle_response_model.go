// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppVulScanCycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppVulScanCycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppVulScanCycleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppVulScanCycleResponseBody) *ModifyAppVulScanCycleResponse
	GetBody() *ModifyAppVulScanCycleResponseBody
}

type ModifyAppVulScanCycleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppVulScanCycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppVulScanCycleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppVulScanCycleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppVulScanCycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppVulScanCycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppVulScanCycleResponse) GetBody() *ModifyAppVulScanCycleResponseBody {
	return s.Body
}

func (s *ModifyAppVulScanCycleResponse) SetHeaders(v map[string]*string) *ModifyAppVulScanCycleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppVulScanCycleResponse) SetStatusCode(v int32) *ModifyAppVulScanCycleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppVulScanCycleResponse) SetBody(v *ModifyAppVulScanCycleResponseBody) *ModifyAppVulScanCycleResponse {
	s.Body = v
	return s
}

func (s *ModifyAppVulScanCycleResponse) Validate() error {
	return dara.Validate(s)
}
