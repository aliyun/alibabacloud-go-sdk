// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgRelationCountAndQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAgRelationCountAndQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAgRelationCountAndQuotaResponse
	GetStatusCode() *int32
	SetBody(v *QueryAgRelationCountAndQuotaResponseBody) *QueryAgRelationCountAndQuotaResponse
	GetBody() *QueryAgRelationCountAndQuotaResponseBody
}

type QueryAgRelationCountAndQuotaResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAgRelationCountAndQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAgRelationCountAndQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAgRelationCountAndQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryAgRelationCountAndQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAgRelationCountAndQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAgRelationCountAndQuotaResponse) GetBody() *QueryAgRelationCountAndQuotaResponseBody {
	return s.Body
}

func (s *QueryAgRelationCountAndQuotaResponse) SetHeaders(v map[string]*string) *QueryAgRelationCountAndQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponse) SetStatusCode(v int32) *QueryAgRelationCountAndQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponse) SetBody(v *QueryAgRelationCountAndQuotaResponseBody) *QueryAgRelationCountAndQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
