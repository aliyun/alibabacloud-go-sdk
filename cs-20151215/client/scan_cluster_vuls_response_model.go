// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanClusterVulsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanClusterVulsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanClusterVulsResponse
	GetStatusCode() *int32
	SetBody(v *ScanClusterVulsResponseBody) *ScanClusterVulsResponse
	GetBody() *ScanClusterVulsResponseBody
}

type ScanClusterVulsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanClusterVulsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanClusterVulsResponse) String() string {
	return dara.Prettify(s)
}

func (s ScanClusterVulsResponse) GoString() string {
	return s.String()
}

func (s *ScanClusterVulsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanClusterVulsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanClusterVulsResponse) GetBody() *ScanClusterVulsResponseBody {
	return s.Body
}

func (s *ScanClusterVulsResponse) SetHeaders(v map[string]*string) *ScanClusterVulsResponse {
	s.Headers = v
	return s
}

func (s *ScanClusterVulsResponse) SetStatusCode(v int32) *ScanClusterVulsResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanClusterVulsResponse) SetBody(v *ScanClusterVulsResponseBody) *ScanClusterVulsResponse {
	s.Body = v
	return s
}

func (s *ScanClusterVulsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
