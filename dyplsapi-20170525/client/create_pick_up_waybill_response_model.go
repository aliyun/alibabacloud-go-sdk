// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePickUpWaybillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePickUpWaybillResponse
	GetStatusCode() *int32
	SetBody(v *CreatePickUpWaybillResponseBody) *CreatePickUpWaybillResponse
	GetBody() *CreatePickUpWaybillResponseBody
}

type CreatePickUpWaybillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePickUpWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePickUpWaybillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillResponse) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePickUpWaybillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePickUpWaybillResponse) GetBody() *CreatePickUpWaybillResponseBody {
	return s.Body
}

func (s *CreatePickUpWaybillResponse) SetHeaders(v map[string]*string) *CreatePickUpWaybillResponse {
	s.Headers = v
	return s
}

func (s *CreatePickUpWaybillResponse) SetStatusCode(v int32) *CreatePickUpWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePickUpWaybillResponse) SetBody(v *CreatePickUpWaybillResponseBody) *CreatePickUpWaybillResponse {
	s.Body = v
	return s
}

func (s *CreatePickUpWaybillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
