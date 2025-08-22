// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnUserRealTimeDeliveryFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnUserRealTimeDeliveryFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnUserRealTimeDeliveryFieldResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) *UpdateDcdnUserRealTimeDeliveryFieldResponse
	GetBody() *UpdateDcdnUserRealTimeDeliveryFieldResponseBody
}

type UpdateDcdnUserRealTimeDeliveryFieldResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnUserRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) GetBody() *UpdateDcdnUserRealTimeDeliveryFieldResponseBody {
	return s.Body
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *UpdateDcdnUserRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) SetStatusCode(v int32) *UpdateDcdnUserRealTimeDeliveryFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) SetBody(v *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) *UpdateDcdnUserRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponse) Validate() error {
	return dara.Validate(s)
}
