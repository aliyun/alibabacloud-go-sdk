// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMajorProtectionBlackIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMajorProtectionBlackIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMajorProtectionBlackIpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMajorProtectionBlackIpResponseBody) *DeleteMajorProtectionBlackIpResponse
	GetBody() *DeleteMajorProtectionBlackIpResponseBody
}

type DeleteMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMajorProtectionBlackIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMajorProtectionBlackIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMajorProtectionBlackIpResponse) GetBody() *DeleteMajorProtectionBlackIpResponseBody {
	return s.Body
}

func (s *DeleteMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *DeleteMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetStatusCode(v int32) *DeleteMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetBody(v *DeleteMajorProtectionBlackIpResponseBody) *DeleteMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
