// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLLogsV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLLogsV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeSQLLogsV2ResponseBody) *DescribeSQLLogsV2Response
	GetBody() *DescribeSQLLogsV2ResponseBody
}

type DescribeSQLLogsV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLLogsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLLogsV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLLogsV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLLogsV2Response) GetBody() *DescribeSQLLogsV2ResponseBody {
	return s.Body
}

func (s *DescribeSQLLogsV2Response) SetHeaders(v map[string]*string) *DescribeSQLLogsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogsV2Response) SetStatusCode(v int32) *DescribeSQLLogsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogsV2Response) SetBody(v *DescribeSQLLogsV2ResponseBody) *DescribeSQLLogsV2Response {
	s.Body = v
	return s
}

func (s *DescribeSQLLogsV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
