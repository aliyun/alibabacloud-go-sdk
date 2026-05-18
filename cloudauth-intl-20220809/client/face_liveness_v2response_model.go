// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceLivenessV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceLivenessV2Response
	GetStatusCode() *int32
	SetBody(v *FaceLivenessV2ResponseBody) *FaceLivenessV2Response
	GetBody() *FaceLivenessV2ResponseBody
}

type FaceLivenessV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceLivenessV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceLivenessV2Response) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2Response) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceLivenessV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceLivenessV2Response) GetBody() *FaceLivenessV2ResponseBody {
	return s.Body
}

func (s *FaceLivenessV2Response) SetHeaders(v map[string]*string) *FaceLivenessV2Response {
	s.Headers = v
	return s
}

func (s *FaceLivenessV2Response) SetStatusCode(v int32) *FaceLivenessV2Response {
	s.StatusCode = &v
	return s
}

func (s *FaceLivenessV2Response) SetBody(v *FaceLivenessV2ResponseBody) *FaceLivenessV2Response {
	s.Body = v
	return s
}

func (s *FaceLivenessV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
