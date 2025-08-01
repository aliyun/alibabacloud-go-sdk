// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotAnalysisResponseBody) *GetHotspotAnalysisResponse
	GetBody() *GetHotspotAnalysisResponseBody
}

type GetHotspotAnalysisResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotAnalysisResponse) GetBody() *GetHotspotAnalysisResponseBody {
	return s.Body
}

func (s *GetHotspotAnalysisResponse) SetHeaders(v map[string]*string) *GetHotspotAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotAnalysisResponse) SetStatusCode(v int32) *GetHotspotAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotAnalysisResponse) SetBody(v *GetHotspotAnalysisResponseBody) *GetHotspotAnalysisResponse {
	s.Body = v
	return s
}

func (s *GetHotspotAnalysisResponse) Validate() error {
	return dara.Validate(s)
}
