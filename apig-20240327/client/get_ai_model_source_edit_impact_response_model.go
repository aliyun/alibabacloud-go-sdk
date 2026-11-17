// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceEditImpactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiModelSourceEditImpactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiModelSourceEditImpactResponse
	GetStatusCode() *int32
	SetBody(v *GetAiModelSourceEditImpactResponseBody) *GetAiModelSourceEditImpactResponse
	GetBody() *GetAiModelSourceEditImpactResponseBody
}

type GetAiModelSourceEditImpactResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiModelSourceEditImpactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiModelSourceEditImpactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceEditImpactResponse) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceEditImpactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiModelSourceEditImpactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiModelSourceEditImpactResponse) GetBody() *GetAiModelSourceEditImpactResponseBody {
	return s.Body
}

func (s *GetAiModelSourceEditImpactResponse) SetHeaders(v map[string]*string) *GetAiModelSourceEditImpactResponse {
	s.Headers = v
	return s
}

func (s *GetAiModelSourceEditImpactResponse) SetStatusCode(v int32) *GetAiModelSourceEditImpactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiModelSourceEditImpactResponse) SetBody(v *GetAiModelSourceEditImpactResponseBody) *GetAiModelSourceEditImpactResponse {
	s.Body = v
	return s
}

func (s *GetAiModelSourceEditImpactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
