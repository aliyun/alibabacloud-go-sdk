// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachCenChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachCenChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DetachCenChildInstanceResponseBody) *DetachCenChildInstanceResponse
	GetBody() *DetachCenChildInstanceResponseBody
}

type DetachCenChildInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachCenChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachCenChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachCenChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachCenChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachCenChildInstanceResponse) GetBody() *DetachCenChildInstanceResponseBody {
	return s.Body
}

func (s *DetachCenChildInstanceResponse) SetHeaders(v map[string]*string) *DetachCenChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DetachCenChildInstanceResponse) SetStatusCode(v int32) *DetachCenChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachCenChildInstanceResponse) SetBody(v *DetachCenChildInstanceResponseBody) *DetachCenChildInstanceResponse {
	s.Body = v
	return s
}

func (s *DetachCenChildInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
