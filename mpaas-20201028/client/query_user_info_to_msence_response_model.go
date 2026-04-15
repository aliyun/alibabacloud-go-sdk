// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoToMsenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserInfoToMsenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserInfoToMsenceResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserInfoToMsenceResponseBody) *QueryUserInfoToMsenceResponse
	GetBody() *QueryUserInfoToMsenceResponseBody
}

type QueryUserInfoToMsenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserInfoToMsenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserInfoToMsenceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoToMsenceResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoToMsenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserInfoToMsenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserInfoToMsenceResponse) GetBody() *QueryUserInfoToMsenceResponseBody {
	return s.Body
}

func (s *QueryUserInfoToMsenceResponse) SetHeaders(v map[string]*string) *QueryUserInfoToMsenceResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoToMsenceResponse) SetStatusCode(v int32) *QueryUserInfoToMsenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoToMsenceResponse) SetBody(v *QueryUserInfoToMsenceResponseBody) *QueryUserInfoToMsenceResponse {
	s.Body = v
	return s
}

func (s *QueryUserInfoToMsenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
