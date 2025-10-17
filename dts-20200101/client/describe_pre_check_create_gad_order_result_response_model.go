// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckCreateGadOrderResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreCheckCreateGadOrderResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreCheckCreateGadOrderResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreCheckCreateGadOrderResultResponseBody) *DescribePreCheckCreateGadOrderResultResponse
	GetBody() *DescribePreCheckCreateGadOrderResultResponseBody
}

type DescribePreCheckCreateGadOrderResultResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreCheckCreateGadOrderResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreCheckCreateGadOrderResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckCreateGadOrderResultResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckCreateGadOrderResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreCheckCreateGadOrderResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreCheckCreateGadOrderResultResponse) GetBody() *DescribePreCheckCreateGadOrderResultResponseBody {
	return s.Body
}

func (s *DescribePreCheckCreateGadOrderResultResponse) SetHeaders(v map[string]*string) *DescribePreCheckCreateGadOrderResultResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponse) SetStatusCode(v int32) *DescribePreCheckCreateGadOrderResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponse) SetBody(v *DescribePreCheckCreateGadOrderResultResponseBody) *DescribePreCheckCreateGadOrderResultResponse {
	s.Body = v
	return s
}

func (s *DescribePreCheckCreateGadOrderResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
