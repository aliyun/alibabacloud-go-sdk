// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsSaleControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnsSaleControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnsSaleControlResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnsSaleControlResponseBody) *CreateEnsSaleControlResponse
	GetBody() *CreateEnsSaleControlResponseBody
}

type CreateEnsSaleControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnsSaleControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnsSaleControlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlResponse) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnsSaleControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnsSaleControlResponse) GetBody() *CreateEnsSaleControlResponseBody {
	return s.Body
}

func (s *CreateEnsSaleControlResponse) SetHeaders(v map[string]*string) *CreateEnsSaleControlResponse {
	s.Headers = v
	return s
}

func (s *CreateEnsSaleControlResponse) SetStatusCode(v int32) *CreateEnsSaleControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnsSaleControlResponse) SetBody(v *CreateEnsSaleControlResponseBody) *CreateEnsSaleControlResponse {
	s.Body = v
	return s
}

func (s *CreateEnsSaleControlResponse) Validate() error {
	return dara.Validate(s)
}
