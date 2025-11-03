// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysForExpressConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagKeysForExpressConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagKeysForExpressConnectResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagKeysForExpressConnectResponseBody) *DescribeTagKeysForExpressConnectResponse
	GetBody() *DescribeTagKeysForExpressConnectResponseBody
}

type DescribeTagKeysForExpressConnectResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagKeysForExpressConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagKeysForExpressConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysForExpressConnectResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysForExpressConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagKeysForExpressConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagKeysForExpressConnectResponse) GetBody() *DescribeTagKeysForExpressConnectResponseBody {
	return s.Body
}

func (s *DescribeTagKeysForExpressConnectResponse) SetHeaders(v map[string]*string) *DescribeTagKeysForExpressConnectResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponse) SetStatusCode(v int32) *DescribeTagKeysForExpressConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponse) SetBody(v *DescribeTagKeysForExpressConnectResponseBody) *DescribeTagKeysForExpressConnectResponse {
	s.Body = v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
