// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearMajorProtectionBlackIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearMajorProtectionBlackIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearMajorProtectionBlackIpResponse
	GetStatusCode() *int32
	SetBody(v *ClearMajorProtectionBlackIpResponseBody) *ClearMajorProtectionBlackIpResponse
	GetBody() *ClearMajorProtectionBlackIpResponseBody
}

type ClearMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearMajorProtectionBlackIpResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearMajorProtectionBlackIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearMajorProtectionBlackIpResponse) GetBody() *ClearMajorProtectionBlackIpResponseBody {
	return s.Body
}

func (s *ClearMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ClearMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ClearMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetBody(v *ClearMajorProtectionBlackIpResponseBody) *ClearMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
