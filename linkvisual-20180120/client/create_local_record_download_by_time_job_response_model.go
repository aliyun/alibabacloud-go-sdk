// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocalRecordDownloadByTimeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLocalRecordDownloadByTimeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLocalRecordDownloadByTimeJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateLocalRecordDownloadByTimeJobResponseBody) *CreateLocalRecordDownloadByTimeJobResponse
	GetBody() *CreateLocalRecordDownloadByTimeJobResponseBody
}

type CreateLocalRecordDownloadByTimeJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocalRecordDownloadByTimeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocalRecordDownloadByTimeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLocalRecordDownloadByTimeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) GetBody() *CreateLocalRecordDownloadByTimeJobResponseBody {
	return s.Body
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetHeaders(v map[string]*string) *CreateLocalRecordDownloadByTimeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetStatusCode(v int32) *CreateLocalRecordDownloadByTimeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) SetBody(v *CreateLocalRecordDownloadByTimeJobResponseBody) *CreateLocalRecordDownloadByTimeJobResponse {
	s.Body = v
	return s
}

func (s *CreateLocalRecordDownloadByTimeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
