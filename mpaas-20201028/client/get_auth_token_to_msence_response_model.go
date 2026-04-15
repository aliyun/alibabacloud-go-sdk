// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenToMsenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthTokenToMsenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthTokenToMsenceResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthTokenToMsenceResponseBody) *GetAuthTokenToMsenceResponse
	GetBody() *GetAuthTokenToMsenceResponseBody
}

type GetAuthTokenToMsenceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthTokenToMsenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthTokenToMsenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenToMsenceResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenToMsenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthTokenToMsenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthTokenToMsenceResponse) GetBody() *GetAuthTokenToMsenceResponseBody {
	return s.Body
}

func (s *GetAuthTokenToMsenceResponse) SetHeaders(v map[string]*string) *GetAuthTokenToMsenceResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenToMsenceResponse) SetStatusCode(v int32) *GetAuthTokenToMsenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenToMsenceResponse) SetBody(v *GetAuthTokenToMsenceResponseBody) *GetAuthTokenToMsenceResponse {
	s.Body = v
	return s
}

func (s *GetAuthTokenToMsenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
