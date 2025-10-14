// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSiteDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSiteDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSiteDeliveryTaskResponseBody) *DeleteSiteDeliveryTaskResponse
	GetBody() *DeleteSiteDeliveryTaskResponseBody
}

type DeleteSiteDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSiteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSiteDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSiteDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSiteDeliveryTaskResponse) GetBody() *DeleteSiteDeliveryTaskResponseBody {
	return s.Body
}

func (s *DeleteSiteDeliveryTaskResponse) SetHeaders(v map[string]*string) *DeleteSiteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteDeliveryTaskResponse) SetStatusCode(v int32) *DeleteSiteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSiteDeliveryTaskResponse) SetBody(v *DeleteSiteDeliveryTaskResponseBody) *DeleteSiteDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteSiteDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
