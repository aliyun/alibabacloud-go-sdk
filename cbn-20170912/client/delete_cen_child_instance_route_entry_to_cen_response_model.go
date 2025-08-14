// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenChildInstanceRouteEntryToCenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenChildInstanceRouteEntryToCenResponseBody) *DeleteCenChildInstanceRouteEntryToCenResponse
	GetBody() *DeleteCenChildInstanceRouteEntryToCenResponseBody
}

type DeleteCenChildInstanceRouteEntryToCenResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenChildInstanceRouteEntryToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToCenResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) GetBody() *DeleteCenChildInstanceRouteEntryToCenResponseBody {
	return s.Body
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToCenResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) SetStatusCode(v int32) *DeleteCenChildInstanceRouteEntryToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) SetBody(v *DeleteCenChildInstanceRouteEntryToCenResponseBody) *DeleteCenChildInstanceRouteEntryToCenResponse {
	s.Body = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToCenResponse) Validate() error {
	return dara.Validate(s)
}
