// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryAnalysisResponseBody) *GetQueryAnalysisResponse
	GetBody() *GetQueryAnalysisResponseBody
}

type GetQueryAnalysisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetQueryAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryAnalysisResponse) GetBody() *GetQueryAnalysisResponseBody {
	return s.Body
}

func (s *GetQueryAnalysisResponse) SetHeaders(v map[string]*string) *GetQueryAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetQueryAnalysisResponse) SetStatusCode(v int32) *GetQueryAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryAnalysisResponse) SetBody(v *GetQueryAnalysisResponseBody) *GetQueryAnalysisResponse {
	s.Body = v
	return s
}

func (s *GetQueryAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
