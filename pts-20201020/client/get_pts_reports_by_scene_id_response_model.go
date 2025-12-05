// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportsBySceneIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsReportsBySceneIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsReportsBySceneIdResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsReportsBySceneIdResponseBody) *GetPtsReportsBySceneIdResponse
	GetBody() *GetPtsReportsBySceneIdResponseBody
}

type GetPtsReportsBySceneIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsReportsBySceneIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsReportsBySceneIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponse) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsReportsBySceneIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsReportsBySceneIdResponse) GetBody() *GetPtsReportsBySceneIdResponseBody {
	return s.Body
}

func (s *GetPtsReportsBySceneIdResponse) SetHeaders(v map[string]*string) *GetPtsReportsBySceneIdResponse {
	s.Headers = v
	return s
}

func (s *GetPtsReportsBySceneIdResponse) SetStatusCode(v int32) *GetPtsReportsBySceneIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponse) SetBody(v *GetPtsReportsBySceneIdResponseBody) *GetPtsReportsBySceneIdResponse {
	s.Body = v
	return s
}

func (s *GetPtsReportsBySceneIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
