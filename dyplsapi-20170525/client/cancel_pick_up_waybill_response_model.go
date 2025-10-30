// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPickUpWaybillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPickUpWaybillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPickUpWaybillResponse
	GetStatusCode() *int32
	SetBody(v *CancelPickUpWaybillResponseBody) *CancelPickUpWaybillResponse
	GetBody() *CancelPickUpWaybillResponseBody
}

type CancelPickUpWaybillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPickUpWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPickUpWaybillResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPickUpWaybillResponse) GoString() string {
	return s.String()
}

func (s *CancelPickUpWaybillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPickUpWaybillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPickUpWaybillResponse) GetBody() *CancelPickUpWaybillResponseBody {
	return s.Body
}

func (s *CancelPickUpWaybillResponse) SetHeaders(v map[string]*string) *CancelPickUpWaybillResponse {
	s.Headers = v
	return s
}

func (s *CancelPickUpWaybillResponse) SetStatusCode(v int32) *CancelPickUpWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPickUpWaybillResponse) SetBody(v *CancelPickUpWaybillResponseBody) *CancelPickUpWaybillResponse {
	s.Body = v
	return s
}

func (s *CancelPickUpWaybillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
