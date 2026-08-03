// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInsightTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInsightTypesResponse
	GetStatusCode() *int32
	SetBody(v *GetInsightTypesResponseBody) *GetInsightTypesResponse
	GetBody() *GetInsightTypesResponseBody
}

type GetInsightTypesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInsightTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInsightTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInsightTypesResponse) GoString() string {
	return s.String()
}

func (s *GetInsightTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInsightTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInsightTypesResponse) GetBody() *GetInsightTypesResponseBody {
	return s.Body
}

func (s *GetInsightTypesResponse) SetHeaders(v map[string]*string) *GetInsightTypesResponse {
	s.Headers = v
	return s
}

func (s *GetInsightTypesResponse) SetStatusCode(v int32) *GetInsightTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInsightTypesResponse) SetBody(v *GetInsightTypesResponseBody) *GetInsightTypesResponse {
	s.Body = v
	return s
}

func (s *GetInsightTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
