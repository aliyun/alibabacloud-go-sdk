// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHcExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHcExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHcExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHcExportInfoResponseBody) *DescribeHcExportInfoResponse
	GetBody() *DescribeHcExportInfoResponseBody
}

type DescribeHcExportInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHcExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHcExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHcExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeHcExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHcExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHcExportInfoResponse) GetBody() *DescribeHcExportInfoResponseBody {
	return s.Body
}

func (s *DescribeHcExportInfoResponse) SetHeaders(v map[string]*string) *DescribeHcExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeHcExportInfoResponse) SetStatusCode(v int32) *DescribeHcExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHcExportInfoResponse) SetBody(v *DescribeHcExportInfoResponseBody) *DescribeHcExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeHcExportInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
