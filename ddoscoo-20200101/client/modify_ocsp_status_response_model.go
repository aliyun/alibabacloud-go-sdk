// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOcspStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOcspStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOcspStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOcspStatusResponseBody) *ModifyOcspStatusResponse
	GetBody() *ModifyOcspStatusResponseBody
}

type ModifyOcspStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOcspStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOcspStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOcspStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyOcspStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOcspStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOcspStatusResponse) GetBody() *ModifyOcspStatusResponseBody {
	return s.Body
}

func (s *ModifyOcspStatusResponse) SetHeaders(v map[string]*string) *ModifyOcspStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyOcspStatusResponse) SetStatusCode(v int32) *ModifyOcspStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOcspStatusResponse) SetBody(v *ModifyOcspStatusResponseBody) *ModifyOcspStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyOcspStatusResponse) Validate() error {
	return dara.Validate(s)
}
