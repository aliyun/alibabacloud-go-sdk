// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePoliciesV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePoliciesV2Response
	GetStatusCode() *int32
	SetBody(v *DescribePoliciesV2ResponseBody) *DescribePoliciesV2Response
	GetBody() *DescribePoliciesV2ResponseBody
}

type DescribePoliciesV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePoliciesV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePoliciesV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2Response) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePoliciesV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePoliciesV2Response) GetBody() *DescribePoliciesV2ResponseBody {
	return s.Body
}

func (s *DescribePoliciesV2Response) SetHeaders(v map[string]*string) *DescribePoliciesV2Response {
	s.Headers = v
	return s
}

func (s *DescribePoliciesV2Response) SetStatusCode(v int32) *DescribePoliciesV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribePoliciesV2Response) SetBody(v *DescribePoliciesV2ResponseBody) *DescribePoliciesV2Response {
	s.Body = v
	return s
}

func (s *DescribePoliciesV2Response) Validate() error {
	return dara.Validate(s)
}
