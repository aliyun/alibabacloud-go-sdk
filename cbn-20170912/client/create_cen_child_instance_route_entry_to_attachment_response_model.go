// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenChildInstanceRouteEntryToAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) *CreateCenChildInstanceRouteEntryToAttachmentResponse
	GetBody() *CreateCenChildInstanceRouteEntryToAttachmentResponseBody
}

type CreateCenChildInstanceRouteEntryToAttachmentResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenChildInstanceRouteEntryToAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) GetBody() *CreateCenChildInstanceRouteEntryToAttachmentResponseBody {
	return s.Body
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) SetHeaders(v map[string]*string) *CreateCenChildInstanceRouteEntryToAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) SetStatusCode(v int32) *CreateCenChildInstanceRouteEntryToAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) SetBody(v *CreateCenChildInstanceRouteEntryToAttachmentResponseBody) *CreateCenChildInstanceRouteEntryToAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
