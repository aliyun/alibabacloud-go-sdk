// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEditingJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEditingJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetEditingJobInfoResponseBody) *GetEditingJobInfoResponse
	GetBody() *GetEditingJobInfoResponseBody
}

type GetEditingJobInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEditingJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEditingJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEditingJobInfoResponse) GoString() string {
	return s.String()
}

func (s *GetEditingJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEditingJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEditingJobInfoResponse) GetBody() *GetEditingJobInfoResponseBody {
	return s.Body
}

func (s *GetEditingJobInfoResponse) SetHeaders(v map[string]*string) *GetEditingJobInfoResponse {
	s.Headers = v
	return s
}

func (s *GetEditingJobInfoResponse) SetStatusCode(v int32) *GetEditingJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEditingJobInfoResponse) SetBody(v *GetEditingJobInfoResponseBody) *GetEditingJobInfoResponse {
	s.Body = v
	return s
}

func (s *GetEditingJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
