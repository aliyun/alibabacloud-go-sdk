// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOnlineTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOnlineTestResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOnlineTestResultResponseBody) *DescribeOnlineTestResultResponse
	GetBody() *DescribeOnlineTestResultResponseBody
}

type DescribeOnlineTestResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnlineTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnlineTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOnlineTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOnlineTestResultResponse) GetBody() *DescribeOnlineTestResultResponseBody {
	return s.Body
}

func (s *DescribeOnlineTestResultResponse) SetHeaders(v map[string]*string) *DescribeOnlineTestResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnlineTestResultResponse) SetStatusCode(v int32) *DescribeOnlineTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnlineTestResultResponse) SetBody(v *DescribeOnlineTestResultResponseBody) *DescribeOnlineTestResultResponse {
	s.Body = v
	return s
}

func (s *DescribeOnlineTestResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
