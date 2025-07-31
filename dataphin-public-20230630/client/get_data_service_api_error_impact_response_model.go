// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiErrorImpactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiErrorImpactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiErrorImpactResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiErrorImpactResponseBody) *GetDataServiceApiErrorImpactResponse
	GetBody() *GetDataServiceApiErrorImpactResponseBody
}

type GetDataServiceApiErrorImpactResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiErrorImpactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiErrorImpactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiErrorImpactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiErrorImpactResponse) GetBody() *GetDataServiceApiErrorImpactResponseBody {
	return s.Body
}

func (s *GetDataServiceApiErrorImpactResponse) SetHeaders(v map[string]*string) *GetDataServiceApiErrorImpactResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiErrorImpactResponse) SetStatusCode(v int32) *GetDataServiceApiErrorImpactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponse) SetBody(v *GetDataServiceApiErrorImpactResponseBody) *GetDataServiceApiErrorImpactResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiErrorImpactResponse) Validate() error {
	return dara.Validate(s)
}
