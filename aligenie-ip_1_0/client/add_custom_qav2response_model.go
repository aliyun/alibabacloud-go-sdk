// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomQAV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomQAV2Response
	GetStatusCode() *int32
	SetBody(v *AddCustomQAV2ResponseBody) *AddCustomQAV2Response
	GetBody() *AddCustomQAV2ResponseBody
}

type AddCustomQAV2Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomQAV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomQAV2Response) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2Response) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomQAV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomQAV2Response) GetBody() *AddCustomQAV2ResponseBody {
	return s.Body
}

func (s *AddCustomQAV2Response) SetHeaders(v map[string]*string) *AddCustomQAV2Response {
	s.Headers = v
	return s
}

func (s *AddCustomQAV2Response) SetStatusCode(v int32) *AddCustomQAV2Response {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAV2Response) SetBody(v *AddCustomQAV2ResponseBody) *AddCustomQAV2Response {
	s.Body = v
	return s
}

func (s *AddCustomQAV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
