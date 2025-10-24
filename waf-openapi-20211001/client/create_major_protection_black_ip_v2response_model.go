// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMajorProtectionBlackIpV2Response
	GetStatusCode() *int32
	SetBody(v *CreateMajorProtectionBlackIpV2ResponseBody) *CreateMajorProtectionBlackIpV2Response
	GetBody() *CreateMajorProtectionBlackIpV2ResponseBody
}

type CreateMajorProtectionBlackIpV2Response struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMajorProtectionBlackIpV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMajorProtectionBlackIpV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpV2Response) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMajorProtectionBlackIpV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMajorProtectionBlackIpV2Response) GetBody() *CreateMajorProtectionBlackIpV2ResponseBody {
	return s.Body
}

func (s *CreateMajorProtectionBlackIpV2Response) SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpV2Response {
	s.Headers = v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Response) SetStatusCode(v int32) *CreateMajorProtectionBlackIpV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Response) SetBody(v *CreateMajorProtectionBlackIpV2ResponseBody) *CreateMajorProtectionBlackIpV2Response {
	s.Body = v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
