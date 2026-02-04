// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenChildInstanceRouteEntriesToAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCenChildInstanceRouteEntriesToAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCenChildInstanceRouteEntriesToAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) *ListCenChildInstanceRouteEntriesToAttachmentResponse
	GetBody() *ListCenChildInstanceRouteEntriesToAttachmentResponseBody
}

type ListCenChildInstanceRouteEntriesToAttachmentResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCenChildInstanceRouteEntriesToAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCenChildInstanceRouteEntriesToAttachmentResponse) GoString() string {
	return s.String()
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) GetBody() *ListCenChildInstanceRouteEntriesToAttachmentResponseBody {
	return s.Body
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) SetHeaders(v map[string]*string) *ListCenChildInstanceRouteEntriesToAttachmentResponse {
	s.Headers = v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) SetStatusCode(v int32) *ListCenChildInstanceRouteEntriesToAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) SetBody(v *ListCenChildInstanceRouteEntriesToAttachmentResponseBody) *ListCenChildInstanceRouteEntriesToAttachmentResponse {
	s.Body = v
	return s
}

func (s *ListCenChildInstanceRouteEntriesToAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
