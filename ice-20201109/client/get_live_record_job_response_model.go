// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveRecordJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveRecordJobResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveRecordJobResponseBody) *GetLiveRecordJobResponse
	GetBody() *GetLiveRecordJobResponseBody
}

type GetLiveRecordJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveRecordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveRecordJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordJobResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRecordJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveRecordJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveRecordJobResponse) GetBody() *GetLiveRecordJobResponseBody {
	return s.Body
}

func (s *GetLiveRecordJobResponse) SetHeaders(v map[string]*string) *GetLiveRecordJobResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRecordJobResponse) SetStatusCode(v int32) *GetLiveRecordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRecordJobResponse) SetBody(v *GetLiveRecordJobResponseBody) *GetLiveRecordJobResponse {
	s.Body = v
	return s
}

func (s *GetLiveRecordJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
