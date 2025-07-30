// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMasterNodeShutDownFailOverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MasterNodeShutDownFailOverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MasterNodeShutDownFailOverResponse
	GetStatusCode() *int32
	SetBody(v *MasterNodeShutDownFailOverResponseBody) *MasterNodeShutDownFailOverResponse
	GetBody() *MasterNodeShutDownFailOverResponseBody
}

type MasterNodeShutDownFailOverResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MasterNodeShutDownFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MasterNodeShutDownFailOverResponse) String() string {
	return dara.Prettify(s)
}

func (s MasterNodeShutDownFailOverResponse) GoString() string {
	return s.String()
}

func (s *MasterNodeShutDownFailOverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MasterNodeShutDownFailOverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MasterNodeShutDownFailOverResponse) GetBody() *MasterNodeShutDownFailOverResponseBody {
	return s.Body
}

func (s *MasterNodeShutDownFailOverResponse) SetHeaders(v map[string]*string) *MasterNodeShutDownFailOverResponse {
	s.Headers = v
	return s
}

func (s *MasterNodeShutDownFailOverResponse) SetStatusCode(v int32) *MasterNodeShutDownFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *MasterNodeShutDownFailOverResponse) SetBody(v *MasterNodeShutDownFailOverResponseBody) *MasterNodeShutDownFailOverResponse {
	s.Body = v
	return s
}

func (s *MasterNodeShutDownFailOverResponse) Validate() error {
	return dara.Validate(s)
}
