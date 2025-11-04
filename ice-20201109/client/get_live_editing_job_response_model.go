// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveEditingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveEditingJobResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveEditingJobResponseBody) *GetLiveEditingJobResponse
	GetBody() *GetLiveEditingJobResponseBody
}

type GetLiveEditingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveEditingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveEditingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingJobResponse) GoString() string {
	return s.String()
}

func (s *GetLiveEditingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveEditingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveEditingJobResponse) GetBody() *GetLiveEditingJobResponseBody {
	return s.Body
}

func (s *GetLiveEditingJobResponse) SetHeaders(v map[string]*string) *GetLiveEditingJobResponse {
	s.Headers = v
	return s
}

func (s *GetLiveEditingJobResponse) SetStatusCode(v int32) *GetLiveEditingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveEditingJobResponse) SetBody(v *GetLiveEditingJobResponseBody) *GetLiveEditingJobResponse {
	s.Body = v
	return s
}

func (s *GetLiveEditingJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
