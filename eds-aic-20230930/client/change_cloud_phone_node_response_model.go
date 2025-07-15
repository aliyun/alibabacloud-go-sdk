// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCloudPhoneNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeCloudPhoneNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeCloudPhoneNodeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeCloudPhoneNodeResponseBody) *ChangeCloudPhoneNodeResponse
	GetBody() *ChangeCloudPhoneNodeResponseBody
}

type ChangeCloudPhoneNodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCloudPhoneNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCloudPhoneNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeCloudPhoneNodeResponse) GoString() string {
	return s.String()
}

func (s *ChangeCloudPhoneNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeCloudPhoneNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeCloudPhoneNodeResponse) GetBody() *ChangeCloudPhoneNodeResponseBody {
	return s.Body
}

func (s *ChangeCloudPhoneNodeResponse) SetHeaders(v map[string]*string) *ChangeCloudPhoneNodeResponse {
	s.Headers = v
	return s
}

func (s *ChangeCloudPhoneNodeResponse) SetStatusCode(v int32) *ChangeCloudPhoneNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCloudPhoneNodeResponse) SetBody(v *ChangeCloudPhoneNodeResponseBody) *ChangeCloudPhoneNodeResponse {
	s.Body = v
	return s
}

func (s *ChangeCloudPhoneNodeResponse) Validate() error {
	return dara.Validate(s)
}
