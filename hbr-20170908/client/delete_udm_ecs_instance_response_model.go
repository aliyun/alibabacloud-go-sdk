// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmEcsInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUdmEcsInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUdmEcsInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUdmEcsInstanceResponseBody) *DeleteUdmEcsInstanceResponse
	GetBody() *DeleteUdmEcsInstanceResponseBody
}

type DeleteUdmEcsInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdmEcsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdmEcsInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmEcsInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdmEcsInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUdmEcsInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUdmEcsInstanceResponse) GetBody() *DeleteUdmEcsInstanceResponseBody {
	return s.Body
}

func (s *DeleteUdmEcsInstanceResponse) SetHeaders(v map[string]*string) *DeleteUdmEcsInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdmEcsInstanceResponse) SetStatusCode(v int32) *DeleteUdmEcsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdmEcsInstanceResponse) SetBody(v *DeleteUdmEcsInstanceResponseBody) *DeleteUdmEcsInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteUdmEcsInstanceResponse) Validate() error {
	return dara.Validate(s)
}
