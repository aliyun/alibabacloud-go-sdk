// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneBaseLineFromReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePtsSceneBaseLineFromReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePtsSceneBaseLineFromReportResponse
	GetStatusCode() *int32
	SetBody(v *CreatePtsSceneBaseLineFromReportResponseBody) *CreatePtsSceneBaseLineFromReportResponse
	GetBody() *CreatePtsSceneBaseLineFromReportResponseBody
}

type CreatePtsSceneBaseLineFromReportResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePtsSceneBaseLineFromReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePtsSceneBaseLineFromReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportResponse) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePtsSceneBaseLineFromReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePtsSceneBaseLineFromReportResponse) GetBody() *CreatePtsSceneBaseLineFromReportResponseBody {
	return s.Body
}

func (s *CreatePtsSceneBaseLineFromReportResponse) SetHeaders(v map[string]*string) *CreatePtsSceneBaseLineFromReportResponse {
	s.Headers = v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponse) SetStatusCode(v int32) *CreatePtsSceneBaseLineFromReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponse) SetBody(v *CreatePtsSceneBaseLineFromReportResponseBody) *CreatePtsSceneBaseLineFromReportResponse {
	s.Body = v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
