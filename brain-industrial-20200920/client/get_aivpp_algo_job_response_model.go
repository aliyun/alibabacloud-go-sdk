// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAivppAlgoJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAivppAlgoJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAivppAlgoJobResponse
	GetStatusCode() *int32
	SetBody(v *GetAivppAlgoJobResponseBody) *GetAivppAlgoJobResponse
	GetBody() *GetAivppAlgoJobResponseBody
}

type GetAivppAlgoJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAivppAlgoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAivppAlgoJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAivppAlgoJobResponse) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAivppAlgoJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAivppAlgoJobResponse) GetBody() *GetAivppAlgoJobResponseBody {
	return s.Body
}

func (s *GetAivppAlgoJobResponse) SetHeaders(v map[string]*string) *GetAivppAlgoJobResponse {
	s.Headers = v
	return s
}

func (s *GetAivppAlgoJobResponse) SetStatusCode(v int32) *GetAivppAlgoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAivppAlgoJobResponse) SetBody(v *GetAivppAlgoJobResponseBody) *GetAivppAlgoJobResponse {
	s.Body = v
	return s
}

func (s *GetAivppAlgoJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
