// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetUserDeliveryTaskResponseBody) *GetUserDeliveryTaskResponse
	GetBody() *GetUserDeliveryTaskResponseBody
}

type GetUserDeliveryTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserDeliveryTaskResponse) GetBody() *GetUserDeliveryTaskResponseBody {
	return s.Body
}

func (s *GetUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *GetUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetUserDeliveryTaskResponse) SetStatusCode(v int32) *GetUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDeliveryTaskResponse) SetBody(v *GetUserDeliveryTaskResponseBody) *GetUserDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *GetUserDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
