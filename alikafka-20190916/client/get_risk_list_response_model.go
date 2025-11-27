// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRiskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRiskListResponse
	GetStatusCode() *int32
	SetBody(v *GetRiskListResponseBody) *GetRiskListResponse
	GetBody() *GetRiskListResponseBody
}

type GetRiskListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRiskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRiskListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRiskListResponse) GoString() string {
	return s.String()
}

func (s *GetRiskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRiskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRiskListResponse) GetBody() *GetRiskListResponseBody {
	return s.Body
}

func (s *GetRiskListResponse) SetHeaders(v map[string]*string) *GetRiskListResponse {
	s.Headers = v
	return s
}

func (s *GetRiskListResponse) SetStatusCode(v int32) *GetRiskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRiskListResponse) SetBody(v *GetRiskListResponseBody) *GetRiskListResponse {
	s.Body = v
	return s
}

func (s *GetRiskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
