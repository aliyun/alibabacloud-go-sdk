// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteDeliveryTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteDeliveryTaskStatusResponseBody) *UpdateSiteDeliveryTaskStatusResponse
	GetBody() *UpdateSiteDeliveryTaskStatusResponseBody
}

type UpdateSiteDeliveryTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteDeliveryTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteDeliveryTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteDeliveryTaskStatusResponse) GetBody() *UpdateSiteDeliveryTaskStatusResponseBody {
	return s.Body
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetStatusCode(v int32) *UpdateSiteDeliveryTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponse) SetBody(v *UpdateSiteDeliveryTaskStatusResponseBody) *UpdateSiteDeliveryTaskStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
