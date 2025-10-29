// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnapshotSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnapshotSettingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnapshotSettingResponseBody) *DescribeSnapshotSettingResponse
	GetBody() *DescribeSnapshotSettingResponseBody
}

type DescribeSnapshotSettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnapshotSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnapshotSettingResponse) GetBody() *DescribeSnapshotSettingResponseBody {
	return s.Body
}

func (s *DescribeSnapshotSettingResponse) SetHeaders(v map[string]*string) *DescribeSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotSettingResponse) SetStatusCode(v int32) *DescribeSnapshotSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotSettingResponse) SetBody(v *DescribeSnapshotSettingResponseBody) *DescribeSnapshotSettingResponse {
	s.Body = v
	return s
}

func (s *DescribeSnapshotSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
