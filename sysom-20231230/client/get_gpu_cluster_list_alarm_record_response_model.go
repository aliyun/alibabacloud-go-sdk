// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterListAlarmRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGpuClusterListAlarmRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGpuClusterListAlarmRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetGpuClusterListAlarmRecordResponseBody) *GetGpuClusterListAlarmRecordResponse
	GetBody() *GetGpuClusterListAlarmRecordResponseBody
}

type GetGpuClusterListAlarmRecordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGpuClusterListAlarmRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGpuClusterListAlarmRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterListAlarmRecordResponse) GoString() string {
	return s.String()
}

func (s *GetGpuClusterListAlarmRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGpuClusterListAlarmRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGpuClusterListAlarmRecordResponse) GetBody() *GetGpuClusterListAlarmRecordResponseBody {
	return s.Body
}

func (s *GetGpuClusterListAlarmRecordResponse) SetHeaders(v map[string]*string) *GetGpuClusterListAlarmRecordResponse {
	s.Headers = v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponse) SetStatusCode(v int32) *GetGpuClusterListAlarmRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponse) SetBody(v *GetGpuClusterListAlarmRecordResponseBody) *GetGpuClusterListAlarmRecordResponse {
	s.Body = v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
