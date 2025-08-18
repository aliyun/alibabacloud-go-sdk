// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSiteDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSiteDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateSiteDeliveryTaskResponseBody) *CreateSiteDeliveryTaskResponse
	GetBody() *CreateSiteDeliveryTaskResponseBody
}

type CreateSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSiteDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSiteDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSiteDeliveryTaskResponse) GetBody() *CreateSiteDeliveryTaskResponseBody {
	return s.Body
}

func (s *CreateSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *CreateSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteDeliveryTaskResponse) SetStatusCode(v int32) *CreateSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSiteDeliveryTaskResponse) SetBody(v *CreateSiteDeliveryTaskResponseBody) *CreateSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateSiteDeliveryTaskResponse) Validate() error {
	return dara.Validate(s)
}
