// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaDNALibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaDNALibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaDNALibResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaDNALibResponseBody) *CreateMediaDNALibResponse
	GetBody() *CreateMediaDNALibResponseBody
}

type CreateMediaDNALibResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaDNALibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaDNALibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaDNALibResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaDNALibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaDNALibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaDNALibResponse) GetBody() *CreateMediaDNALibResponseBody {
	return s.Body
}

func (s *CreateMediaDNALibResponse) SetHeaders(v map[string]*string) *CreateMediaDNALibResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaDNALibResponse) SetStatusCode(v int32) *CreateMediaDNALibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaDNALibResponse) SetBody(v *CreateMediaDNALibResponseBody) *CreateMediaDNALibResponse {
	s.Body = v
	return s
}

func (s *CreateMediaDNALibResponse) Validate() error {
	return dara.Validate(s)
}
