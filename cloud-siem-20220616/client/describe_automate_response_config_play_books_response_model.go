// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigPlayBooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigPlayBooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutomateResponseConfigPlayBooksResponseBody) *DescribeAutomateResponseConfigPlayBooksResponse
	GetBody() *DescribeAutomateResponseConfigPlayBooksResponseBody
}

type DescribeAutomateResponseConfigPlayBooksResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigPlayBooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) GetBody() *DescribeAutomateResponseConfigPlayBooksResponseBody {
	return s.Body
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetBody(v *DescribeAutomateResponseConfigPlayBooksResponseBody) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Body = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) Validate() error {
	return dara.Validate(s)
}
