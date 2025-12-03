// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectInfoResponseBody) *GetProjectInfoResponse
	GetBody() *GetProjectInfoResponseBody
}

type GetProjectInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectInfoResponse) GetBody() *GetProjectInfoResponseBody {
	return s.Body
}

func (s *GetProjectInfoResponse) SetHeaders(v map[string]*string) *GetProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetProjectInfoResponse) SetStatusCode(v int32) *GetProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectInfoResponse) SetBody(v *GetProjectInfoResponseBody) *GetProjectInfoResponse {
	s.Body = v
	return s
}

func (s *GetProjectInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
