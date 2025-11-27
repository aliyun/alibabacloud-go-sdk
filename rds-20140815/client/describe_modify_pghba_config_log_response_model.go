// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyPGHbaConfigLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModifyPGHbaConfigLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModifyPGHbaConfigLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModifyPGHbaConfigLogResponseBody) *DescribeModifyPGHbaConfigLogResponse
	GetBody() *DescribeModifyPGHbaConfigLogResponseBody
}

type DescribeModifyPGHbaConfigLogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModifyPGHbaConfigLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModifyPGHbaConfigLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModifyPGHbaConfigLogResponse) GetBody() *DescribeModifyPGHbaConfigLogResponseBody {
	return s.Body
}

func (s *DescribeModifyPGHbaConfigLogResponse) SetHeaders(v map[string]*string) *DescribeModifyPGHbaConfigLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponse) SetStatusCode(v int32) *DescribeModifyPGHbaConfigLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponse) SetBody(v *DescribeModifyPGHbaConfigLogResponseBody) *DescribeModifyPGHbaConfigLogResponse {
	s.Body = v
	return s
}

func (s *DescribeModifyPGHbaConfigLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
