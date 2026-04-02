// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FaceCompareV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FaceCompareV2Response
	GetStatusCode() *int32
	SetBody(v *FaceCompareV2ResponseBody) *FaceCompareV2Response
	GetBody() *FaceCompareV2ResponseBody
}

type FaceCompareV2Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FaceCompareV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FaceCompareV2Response) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2Response) GoString() string {
	return s.String()
}

func (s *FaceCompareV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FaceCompareV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FaceCompareV2Response) GetBody() *FaceCompareV2ResponseBody {
	return s.Body
}

func (s *FaceCompareV2Response) SetHeaders(v map[string]*string) *FaceCompareV2Response {
	s.Headers = v
	return s
}

func (s *FaceCompareV2Response) SetStatusCode(v int32) *FaceCompareV2Response {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareV2Response) SetBody(v *FaceCompareV2ResponseBody) *FaceCompareV2Response {
	s.Body = v
	return s
}

func (s *FaceCompareV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
