// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADConnectorOfficeSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyADConnectorOfficeSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyADConnectorOfficeSiteResponse
	GetStatusCode() *int32
	SetBody(v *ModifyADConnectorOfficeSiteResponseBody) *ModifyADConnectorOfficeSiteResponse
	GetBody() *ModifyADConnectorOfficeSiteResponseBody
}

type ModifyADConnectorOfficeSiteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyADConnectorOfficeSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyADConnectorOfficeSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyADConnectorOfficeSiteResponse) GetBody() *ModifyADConnectorOfficeSiteResponseBody {
	return s.Body
}

func (s *ModifyADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *ModifyADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponse) SetStatusCode(v int32) *ModifyADConnectorOfficeSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponse) SetBody(v *ModifyADConnectorOfficeSiteResponseBody) *ModifyADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponse) Validate() error {
	return dara.Validate(s)
}
