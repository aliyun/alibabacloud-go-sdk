// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnsRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnsRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnsRouteEntryResponseBody) *CreateEnsRouteEntryResponse
	GetBody() *CreateEnsRouteEntryResponseBody
}

type CreateEnsRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnsRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnsRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateEnsRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnsRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnsRouteEntryResponse) GetBody() *CreateEnsRouteEntryResponseBody {
	return s.Body
}

func (s *CreateEnsRouteEntryResponse) SetHeaders(v map[string]*string) *CreateEnsRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateEnsRouteEntryResponse) SetStatusCode(v int32) *CreateEnsRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnsRouteEntryResponse) SetBody(v *CreateEnsRouteEntryResponseBody) *CreateEnsRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateEnsRouteEntryResponse) Validate() error {
	return dara.Validate(s)
}
