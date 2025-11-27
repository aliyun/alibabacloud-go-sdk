// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionPrecheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpgradeMajorVersionPrecheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) *DescribeUpgradeMajorVersionPrecheckTaskResponse
	GetBody() *DescribeUpgradeMajorVersionPrecheckTaskResponseBody
}

type DescribeUpgradeMajorVersionPrecheckTaskResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpgradeMajorVersionPrecheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionPrecheckTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) GetBody() *DescribeUpgradeMajorVersionPrecheckTaskResponseBody {
	return s.Body
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) SetHeaders(v map[string]*string) *DescribeUpgradeMajorVersionPrecheckTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) SetStatusCode(v int32) *DescribeUpgradeMajorVersionPrecheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) SetBody(v *DescribeUpgradeMajorVersionPrecheckTaskResponseBody) *DescribeUpgradeMajorVersionPrecheckTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeUpgradeMajorVersionPrecheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
