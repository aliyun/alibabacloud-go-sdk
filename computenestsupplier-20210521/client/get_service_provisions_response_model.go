// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvisionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceProvisionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceProvisionsResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceProvisionsResponseBody) *GetServiceProvisionsResponse
	GetBody() *GetServiceProvisionsResponseBody
}

type GetServiceProvisionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceProvisionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceProvisionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceProvisionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceProvisionsResponse) GetBody() *GetServiceProvisionsResponseBody {
	return s.Body
}

func (s *GetServiceProvisionsResponse) SetHeaders(v map[string]*string) *GetServiceProvisionsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceProvisionsResponse) SetStatusCode(v int32) *GetServiceProvisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceProvisionsResponse) SetBody(v *GetServiceProvisionsResponseBody) *GetServiceProvisionsResponse {
	s.Body = v
	return s
}

func (s *GetServiceProvisionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
