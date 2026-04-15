// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPayOrderToMsenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPayOrderToMsenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPayOrderToMsenceResponse
	GetStatusCode() *int32
	SetBody(v *QueryPayOrderToMsenceResponseBody) *QueryPayOrderToMsenceResponse
	GetBody() *QueryPayOrderToMsenceResponseBody
}

type QueryPayOrderToMsenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPayOrderToMsenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPayOrderToMsenceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPayOrderToMsenceResponse) GoString() string {
	return s.String()
}

func (s *QueryPayOrderToMsenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPayOrderToMsenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPayOrderToMsenceResponse) GetBody() *QueryPayOrderToMsenceResponseBody {
	return s.Body
}

func (s *QueryPayOrderToMsenceResponse) SetHeaders(v map[string]*string) *QueryPayOrderToMsenceResponse {
	s.Headers = v
	return s
}

func (s *QueryPayOrderToMsenceResponse) SetStatusCode(v int32) *QueryPayOrderToMsenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPayOrderToMsenceResponse) SetBody(v *QueryPayOrderToMsenceResponseBody) *QueryPayOrderToMsenceResponse {
	s.Body = v
	return s
}

func (s *QueryPayOrderToMsenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
