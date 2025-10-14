// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSiteDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSiteDeliveryTaskResponseBody) *UpdateSiteDeliveryTaskResponse
	GetBody() *UpdateSiteDeliveryTaskResponseBody
}

type UpdateSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSiteDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSiteDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSiteDeliveryTaskResponse) GetBody() *UpdateSiteDeliveryTaskResponseBody {
	return s.Body
}

func (s *UpdateSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *UpdateSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateSiteDeliveryTaskResponse) SetStatusCode(v int32) *UpdateSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSiteDeliveryTaskResponse) SetBody(v *UpdateSiteDeliveryTaskResponseBody) *UpdateSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateSiteDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
