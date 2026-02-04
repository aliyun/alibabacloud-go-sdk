// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenChildInstanceRouteEntryToAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) *DeleteCenChildInstanceRouteEntryToAttachmentResponse
	GetBody() *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody
}

type DeleteCenChildInstanceRouteEntryToAttachmentResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) GetBody() *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody {
	return s.Body
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) SetHeaders(v map[string]*string) *DeleteCenChildInstanceRouteEntryToAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) SetStatusCode(v int32) *DeleteCenChildInstanceRouteEntryToAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) SetBody(v *DeleteCenChildInstanceRouteEntryToAttachmentResponseBody) *DeleteCenChildInstanceRouteEntryToAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
