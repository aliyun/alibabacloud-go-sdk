// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecallManagementJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecallManagementJobResponse
	GetStatusCode() *int32
	SetBody(v *GetRecallManagementJobResponseBody) *GetRecallManagementJobResponse
	GetBody() *GetRecallManagementJobResponseBody
}

type GetRecallManagementJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecallManagementJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecallManagementJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementJobResponse) GoString() string {
	return s.String()
}

func (s *GetRecallManagementJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecallManagementJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecallManagementJobResponse) GetBody() *GetRecallManagementJobResponseBody {
	return s.Body
}

func (s *GetRecallManagementJobResponse) SetHeaders(v map[string]*string) *GetRecallManagementJobResponse {
	s.Headers = v
	return s
}

func (s *GetRecallManagementJobResponse) SetStatusCode(v int32) *GetRecallManagementJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecallManagementJobResponse) SetBody(v *GetRecallManagementJobResponseBody) *GetRecallManagementJobResponse {
	s.Body = v
	return s
}

func (s *GetRecallManagementJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
