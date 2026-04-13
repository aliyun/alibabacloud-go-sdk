// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinlogFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBinlogFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBinlogFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBinlogFilesResponseBody) *DescribeBinlogFilesResponse
	GetBody() *DescribeBinlogFilesResponseBody
}

type DescribeBinlogFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBinlogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBinlogFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBinlogFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBinlogFilesResponse) GetBody() *DescribeBinlogFilesResponseBody {
	return s.Body
}

func (s *DescribeBinlogFilesResponse) SetHeaders(v map[string]*string) *DescribeBinlogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBinlogFilesResponse) SetStatusCode(v int32) *DescribeBinlogFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBinlogFilesResponse) SetBody(v *DescribeBinlogFilesResponseBody) *DescribeBinlogFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeBinlogFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
