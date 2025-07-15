// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataChannelCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataChannelCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataChannelCredentialsResponseBody) *GetDataChannelCredentialsResponse
	GetBody() *GetDataChannelCredentialsResponseBody
}

type GetDataChannelCredentialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataChannelCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataChannelCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataChannelCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataChannelCredentialsResponse) GetBody() *GetDataChannelCredentialsResponseBody {
	return s.Body
}

func (s *GetDataChannelCredentialsResponse) SetHeaders(v map[string]*string) *GetDataChannelCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetDataChannelCredentialsResponse) SetStatusCode(v int32) *GetDataChannelCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataChannelCredentialsResponse) SetBody(v *GetDataChannelCredentialsResponseBody) *GetDataChannelCredentialsResponse {
	s.Body = v
	return s
}

func (s *GetDataChannelCredentialsResponse) Validate() error {
	return dara.Validate(s)
}
