// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpRiskDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpRiskDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpRiskDataResponse
	GetStatusCode() *int32
	SetBody(v *GetOpRiskDataResponseBody) *GetOpRiskDataResponse
	GetBody() *GetOpRiskDataResponseBody
}

type GetOpRiskDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpRiskDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpRiskDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpRiskDataResponse) GoString() string {
	return s.String()
}

func (s *GetOpRiskDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpRiskDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpRiskDataResponse) GetBody() *GetOpRiskDataResponseBody {
	return s.Body
}

func (s *GetOpRiskDataResponse) SetHeaders(v map[string]*string) *GetOpRiskDataResponse {
	s.Headers = v
	return s
}

func (s *GetOpRiskDataResponse) SetStatusCode(v int32) *GetOpRiskDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpRiskDataResponse) SetBody(v *GetOpRiskDataResponseBody) *GetOpRiskDataResponse {
	s.Body = v
	return s
}

func (s *GetOpRiskDataResponse) Validate() error {
	return dara.Validate(s)
}
