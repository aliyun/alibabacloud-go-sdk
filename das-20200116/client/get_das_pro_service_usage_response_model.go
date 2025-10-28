// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasProServiceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDasProServiceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDasProServiceUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetDasProServiceUsageResponseBody) *GetDasProServiceUsageResponse
	GetBody() *GetDasProServiceUsageResponseBody
}

type GetDasProServiceUsageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDasProServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDasProServiceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDasProServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDasProServiceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDasProServiceUsageResponse) GetBody() *GetDasProServiceUsageResponseBody {
	return s.Body
}

func (s *GetDasProServiceUsageResponse) SetHeaders(v map[string]*string) *GetDasProServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetDasProServiceUsageResponse) SetStatusCode(v int32) *GetDasProServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDasProServiceUsageResponse) SetBody(v *GetDasProServiceUsageResponseBody) *GetDasProServiceUsageResponse {
	s.Body = v
	return s
}

func (s *GetDasProServiceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
