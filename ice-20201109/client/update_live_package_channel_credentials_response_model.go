// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePackageChannelCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePackageChannelCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePackageChannelCredentialsResponseBody) *UpdateLivePackageChannelCredentialsResponse
	GetBody() *UpdateLivePackageChannelCredentialsResponseBody
}

type UpdateLivePackageChannelCredentialsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePackageChannelCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePackageChannelCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelCredentialsResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePackageChannelCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePackageChannelCredentialsResponse) GetBody() *UpdateLivePackageChannelCredentialsResponseBody {
	return s.Body
}

func (s *UpdateLivePackageChannelCredentialsResponse) SetHeaders(v map[string]*string) *UpdateLivePackageChannelCredentialsResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponse) SetStatusCode(v int32) *UpdateLivePackageChannelCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponse) SetBody(v *UpdateLivePackageChannelCredentialsResponseBody) *UpdateLivePackageChannelCredentialsResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponse) Validate() error {
	return dara.Validate(s)
}
