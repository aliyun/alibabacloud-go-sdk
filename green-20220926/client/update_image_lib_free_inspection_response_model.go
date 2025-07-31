// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibFreeInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageLibFreeInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageLibFreeInspectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageLibFreeInspectionResponseBody) *UpdateImageLibFreeInspectionResponse
	GetBody() *UpdateImageLibFreeInspectionResponseBody
}

type UpdateImageLibFreeInspectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageLibFreeInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageLibFreeInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibFreeInspectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageLibFreeInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageLibFreeInspectionResponse) GetBody() *UpdateImageLibFreeInspectionResponseBody {
	return s.Body
}

func (s *UpdateImageLibFreeInspectionResponse) SetHeaders(v map[string]*string) *UpdateImageLibFreeInspectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageLibFreeInspectionResponse) SetStatusCode(v int32) *UpdateImageLibFreeInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponse) SetBody(v *UpdateImageLibFreeInspectionResponseBody) *UpdateImageLibFreeInspectionResponse {
	s.Body = v
	return s
}

func (s *UpdateImageLibFreeInspectionResponse) Validate() error {
	return dara.Validate(s)
}
