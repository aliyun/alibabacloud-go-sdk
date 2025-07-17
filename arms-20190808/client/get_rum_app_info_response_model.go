// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRumAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRumAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetRumAppInfoResponseBody) *GetRumAppInfoResponse
	GetBody() *GetRumAppInfoResponseBody
}

type GetRumAppInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRumAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRumAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRumAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRumAppInfoResponse) GetBody() *GetRumAppInfoResponseBody {
	return s.Body
}

func (s *GetRumAppInfoResponse) SetHeaders(v map[string]*string) *GetRumAppInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRumAppInfoResponse) SetStatusCode(v int32) *GetRumAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRumAppInfoResponse) SetBody(v *GetRumAppInfoResponseBody) *GetRumAppInfoResponse {
	s.Body = v
	return s
}

func (s *GetRumAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
