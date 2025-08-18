// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteDeliveryTaskResponseBody) *GetSiteDeliveryTaskResponse
	GetBody() *GetSiteDeliveryTaskResponseBody
}

type GetSiteDeliveryTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteDeliveryTaskResponse) GetBody() *GetSiteDeliveryTaskResponseBody {
	return s.Body
}

func (s *GetSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *GetSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSiteDeliveryTaskResponse) SetStatusCode(v int32) *GetSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteDeliveryTaskResponse) SetBody(v *GetSiteDeliveryTaskResponseBody) *GetSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *GetSiteDeliveryTaskResponse) Validate() error {
	return dara.Validate(s)
}
