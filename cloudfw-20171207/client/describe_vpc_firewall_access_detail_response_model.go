// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAccessDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcFirewallAccessDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcFirewallAccessDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcFirewallAccessDetailResponseBody) *DescribeVpcFirewallAccessDetailResponse
	GetBody() *DescribeVpcFirewallAccessDetailResponseBody
}

type DescribeVpcFirewallAccessDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallAccessDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallAccessDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAccessDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAccessDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcFirewallAccessDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcFirewallAccessDetailResponse) GetBody() *DescribeVpcFirewallAccessDetailResponseBody {
	return s.Body
}

func (s *DescribeVpcFirewallAccessDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAccessDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallAccessDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponse) SetBody(v *DescribeVpcFirewallAccessDetailResponseBody) *DescribeVpcFirewallAccessDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
