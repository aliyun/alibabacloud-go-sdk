// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackUpExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackUpExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackUpExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackUpExportInfoResponseBody) *DescribeBackUpExportInfoResponse
	GetBody() *DescribeBackUpExportInfoResponseBody
}

type DescribeBackUpExportInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackUpExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackUpExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackUpExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackUpExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackUpExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackUpExportInfoResponse) GetBody() *DescribeBackUpExportInfoResponseBody {
	return s.Body
}

func (s *DescribeBackUpExportInfoResponse) SetHeaders(v map[string]*string) *DescribeBackUpExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackUpExportInfoResponse) SetStatusCode(v int32) *DescribeBackUpExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackUpExportInfoResponse) SetBody(v *DescribeBackUpExportInfoResponseBody) *DescribeBackUpExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeBackUpExportInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
