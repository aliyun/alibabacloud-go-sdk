// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStructureImportTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStructureImportTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStructureImportTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStructureImportTaskInfoResponseBody) *DescribeStructureImportTaskInfoResponse
	GetBody() *DescribeStructureImportTaskInfoResponseBody
}

type DescribeStructureImportTaskInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStructureImportTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStructureImportTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStructureImportTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeStructureImportTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStructureImportTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStructureImportTaskInfoResponse) GetBody() *DescribeStructureImportTaskInfoResponseBody {
	return s.Body
}

func (s *DescribeStructureImportTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeStructureImportTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeStructureImportTaskInfoResponse) SetStatusCode(v int32) *DescribeStructureImportTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStructureImportTaskInfoResponse) SetBody(v *DescribeStructureImportTaskInfoResponseBody) *DescribeStructureImportTaskInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeStructureImportTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
