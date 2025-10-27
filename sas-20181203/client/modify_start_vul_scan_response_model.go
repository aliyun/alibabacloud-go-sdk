// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStartVulScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStartVulScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStartVulScanResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStartVulScanResponseBody) *ModifyStartVulScanResponse
	GetBody() *ModifyStartVulScanResponseBody
}

type ModifyStartVulScanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStartVulScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStartVulScanResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStartVulScanResponse) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStartVulScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStartVulScanResponse) GetBody() *ModifyStartVulScanResponseBody {
	return s.Body
}

func (s *ModifyStartVulScanResponse) SetHeaders(v map[string]*string) *ModifyStartVulScanResponse {
	s.Headers = v
	return s
}

func (s *ModifyStartVulScanResponse) SetStatusCode(v int32) *ModifyStartVulScanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStartVulScanResponse) SetBody(v *ModifyStartVulScanResponseBody) *ModifyStartVulScanResponse {
	s.Body = v
	return s
}

func (s *ModifyStartVulScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
