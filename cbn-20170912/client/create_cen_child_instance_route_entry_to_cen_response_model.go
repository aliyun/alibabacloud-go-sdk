// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenChildInstanceRouteEntryToCenResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenChildInstanceRouteEntryToCenResponseBody) *CreateCenChildInstanceRouteEntryToCenResponse
	GetBody() *CreateCenChildInstanceRouteEntryToCenResponseBody
}

type CreateCenChildInstanceRouteEntryToCenResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenChildInstanceRouteEntryToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToCenResponse) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) GetBody() *CreateCenChildInstanceRouteEntryToCenResponseBody {
	return s.Body
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToCenResponse {
	s.Headers = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) SetStatusCode(v int32) *CreateCenChildInstanceRouteEntryToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) SetBody(v *CreateCenChildInstanceRouteEntryToCenResponseBody) *CreateCenChildInstanceRouteEntryToCenResponse {
	s.Body = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToCenResponse) Validate() error {
	return dara.Validate(s)
}
