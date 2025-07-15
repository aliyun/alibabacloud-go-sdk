// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPhoneNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudPhoneNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudPhoneNodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudPhoneNodeResponseBody) *CreateCloudPhoneNodeResponse
	GetBody() *CreateCloudPhoneNodeResponseBody
}

type CreateCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudPhoneNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudPhoneNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudPhoneNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudPhoneNodeResponse) GetBody() *CreateCloudPhoneNodeResponseBody {
	return s.Body
}

func (s *CreateCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *CreateCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudPhoneNodeResponse) SetStatusCode(v int32) *CreateCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudPhoneNodeResponse) SetBody(v *CreateCloudPhoneNodeResponseBody) *CreateCloudPhoneNodeResponse {
	s.Body = v
	return s
}

func (s *CreateCloudPhoneNodeResponse) Validate() error {
	return dara.Validate(s)
}
