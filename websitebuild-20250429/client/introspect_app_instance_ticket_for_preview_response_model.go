// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntrospectAppInstanceTicketForPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntrospectAppInstanceTicketForPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntrospectAppInstanceTicketForPreviewResponse
	GetStatusCode() *int32
	SetBody(v *IntrospectAppInstanceTicketForPreviewResponseBody) *IntrospectAppInstanceTicketForPreviewResponse
	GetBody() *IntrospectAppInstanceTicketForPreviewResponseBody
}

type IntrospectAppInstanceTicketForPreviewResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntrospectAppInstanceTicketForPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntrospectAppInstanceTicketForPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s IntrospectAppInstanceTicketForPreviewResponse) GoString() string {
	return s.String()
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) GetBody() *IntrospectAppInstanceTicketForPreviewResponseBody {
	return s.Body
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) SetHeaders(v map[string]*string) *IntrospectAppInstanceTicketForPreviewResponse {
	s.Headers = v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) SetStatusCode(v int32) *IntrospectAppInstanceTicketForPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) SetBody(v *IntrospectAppInstanceTicketForPreviewResponseBody) *IntrospectAppInstanceTicketForPreviewResponse {
	s.Body = v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
