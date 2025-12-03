// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSprintInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSprintInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSprintInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSprintInfoResponseBody) *GetSprintInfoResponse
	GetBody() *GetSprintInfoResponseBody
}

type GetSprintInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSprintInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSprintInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSprintInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSprintInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSprintInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSprintInfoResponse) GetBody() *GetSprintInfoResponseBody {
	return s.Body
}

func (s *GetSprintInfoResponse) SetHeaders(v map[string]*string) *GetSprintInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSprintInfoResponse) SetStatusCode(v int32) *GetSprintInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSprintInfoResponse) SetBody(v *GetSprintInfoResponseBody) *GetSprintInfoResponse {
	s.Body = v
	return s
}

func (s *GetSprintInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
