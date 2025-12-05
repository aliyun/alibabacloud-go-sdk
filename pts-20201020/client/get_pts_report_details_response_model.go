// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsReportDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsReportDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsReportDetailsResponseBody) *GetPtsReportDetailsResponse
	GetBody() *GetPtsReportDetailsResponseBody
}

type GetPtsReportDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsReportDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsReportDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsReportDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsReportDetailsResponse) GetBody() *GetPtsReportDetailsResponseBody {
	return s.Body
}

func (s *GetPtsReportDetailsResponse) SetHeaders(v map[string]*string) *GetPtsReportDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetPtsReportDetailsResponse) SetStatusCode(v int32) *GetPtsReportDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsReportDetailsResponse) SetBody(v *GetPtsReportDetailsResponseBody) *GetPtsReportDetailsResponse {
	s.Body = v
	return s
}

func (s *GetPtsReportDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
