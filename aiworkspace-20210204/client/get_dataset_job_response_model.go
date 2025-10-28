// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetJobResponseBody) *GetDatasetJobResponse
	GetBody() *GetDatasetJobResponseBody
}

type GetDatasetJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetJobResponse) GetBody() *GetDatasetJobResponseBody {
	return s.Body
}

func (s *GetDatasetJobResponse) SetHeaders(v map[string]*string) *GetDatasetJobResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetJobResponse) SetStatusCode(v int32) *GetDatasetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetJobResponse) SetBody(v *GetDatasetJobResponseBody) *GetDatasetJobResponse {
	s.Body = v
	return s
}

func (s *GetDatasetJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
