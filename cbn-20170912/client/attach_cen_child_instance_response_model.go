// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachCenChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachCenChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AttachCenChildInstanceResponseBody) *AttachCenChildInstanceResponse
	GetBody() *AttachCenChildInstanceResponseBody
}

type AttachCenChildInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachCenChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachCenChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachCenChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachCenChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachCenChildInstanceResponse) GetBody() *AttachCenChildInstanceResponseBody {
	return s.Body
}

func (s *AttachCenChildInstanceResponse) SetHeaders(v map[string]*string) *AttachCenChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachCenChildInstanceResponse) SetStatusCode(v int32) *AttachCenChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachCenChildInstanceResponse) SetBody(v *AttachCenChildInstanceResponseBody) *AttachCenChildInstanceResponse {
	s.Body = v
	return s
}

func (s *AttachCenChildInstanceResponse) Validate() error {
	return dara.Validate(s)
}
