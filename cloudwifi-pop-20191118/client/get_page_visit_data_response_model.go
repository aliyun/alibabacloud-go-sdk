// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageVisitDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageVisitDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageVisitDataResponse
	GetStatusCode() *int32
	SetBody(v *GetPageVisitDataResponseBody) *GetPageVisitDataResponse
	GetBody() *GetPageVisitDataResponseBody
}

type GetPageVisitDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageVisitDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageVisitDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageVisitDataResponse) GoString() string {
	return s.String()
}

func (s *GetPageVisitDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageVisitDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageVisitDataResponse) GetBody() *GetPageVisitDataResponseBody {
	return s.Body
}

func (s *GetPageVisitDataResponse) SetHeaders(v map[string]*string) *GetPageVisitDataResponse {
	s.Headers = v
	return s
}

func (s *GetPageVisitDataResponse) SetStatusCode(v int32) *GetPageVisitDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageVisitDataResponse) SetBody(v *GetPageVisitDataResponseBody) *GetPageVisitDataResponse {
	s.Body = v
	return s
}

func (s *GetPageVisitDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
