// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceWithEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertSourceWithEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertSourceWithEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertSourceWithEventResponseBody) *DescribeAlertSourceWithEventResponse
	GetBody() *DescribeAlertSourceWithEventResponseBody
}

type DescribeAlertSourceWithEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSourceWithEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSourceWithEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertSourceWithEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertSourceWithEventResponse) GetBody() *DescribeAlertSourceWithEventResponseBody {
	return s.Body
}

func (s *DescribeAlertSourceWithEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceWithEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetStatusCode(v int32) *DescribeAlertSourceWithEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetBody(v *DescribeAlertSourceWithEventResponseBody) *DescribeAlertSourceWithEventResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
