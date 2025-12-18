// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCompliancePacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyCompliancePacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyCompliancePacksResponse
	GetStatusCode() *int32
	SetBody(v *CopyCompliancePacksResponseBody) *CopyCompliancePacksResponse
	GetBody() *CopyCompliancePacksResponseBody
}

type CopyCompliancePacksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyCompliancePacksResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *CopyCompliancePacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyCompliancePacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyCompliancePacksResponse) GetBody() *CopyCompliancePacksResponseBody {
	return s.Body
}

func (s *CopyCompliancePacksResponse) SetHeaders(v map[string]*string) *CopyCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *CopyCompliancePacksResponse) SetStatusCode(v int32) *CopyCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyCompliancePacksResponse) SetBody(v *CopyCompliancePacksResponseBody) *CopyCompliancePacksResponse {
	s.Body = v
	return s
}

func (s *CopyCompliancePacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
