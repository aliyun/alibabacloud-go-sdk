// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordDownloadByTimeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecordDownloadByTimeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecordDownloadByTimeJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecordDownloadByTimeJobResponseBody) *CreateRecordDownloadByTimeJobResponse
	GetBody() *CreateRecordDownloadByTimeJobResponseBody
}

type CreateRecordDownloadByTimeJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecordDownloadByTimeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecordDownloadByTimeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordDownloadByTimeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordDownloadByTimeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecordDownloadByTimeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecordDownloadByTimeJobResponse) GetBody() *CreateRecordDownloadByTimeJobResponseBody {
	return s.Body
}

func (s *CreateRecordDownloadByTimeJobResponse) SetHeaders(v map[string]*string) *CreateRecordDownloadByTimeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponse) SetStatusCode(v int32) *CreateRecordDownloadByTimeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponse) SetBody(v *CreateRecordDownloadByTimeJobResponseBody) *CreateRecordDownloadByTimeJobResponse {
	s.Body = v
	return s
}

func (s *CreateRecordDownloadByTimeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
