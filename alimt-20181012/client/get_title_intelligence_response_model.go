// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleIntelligenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTitleIntelligenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTitleIntelligenceResponse
	GetStatusCode() *int32
	SetBody(v *GetTitleIntelligenceResponseBody) *GetTitleIntelligenceResponse
	GetBody() *GetTitleIntelligenceResponseBody
}

type GetTitleIntelligenceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleIntelligenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTitleIntelligenceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTitleIntelligenceResponse) GoString() string {
	return s.String()
}

func (s *GetTitleIntelligenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTitleIntelligenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTitleIntelligenceResponse) GetBody() *GetTitleIntelligenceResponseBody {
	return s.Body
}

func (s *GetTitleIntelligenceResponse) SetHeaders(v map[string]*string) *GetTitleIntelligenceResponse {
	s.Headers = v
	return s
}

func (s *GetTitleIntelligenceResponse) SetStatusCode(v int32) *GetTitleIntelligenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleIntelligenceResponse) SetBody(v *GetTitleIntelligenceResponseBody) *GetTitleIntelligenceResponse {
	s.Body = v
	return s
}

func (s *GetTitleIntelligenceResponse) Validate() error {
	return dara.Validate(s)
}
