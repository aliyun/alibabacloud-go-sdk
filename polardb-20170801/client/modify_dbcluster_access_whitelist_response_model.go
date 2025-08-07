// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterAccessWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterAccessWhitelistResponseBody) *ModifyDBClusterAccessWhitelistResponse
	GetBody() *ModifyDBClusterAccessWhitelistResponseBody
}

type ModifyDBClusterAccessWhitelistResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterAccessWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterAccessWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterAccessWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterAccessWhitelistResponse) GetBody() *ModifyDBClusterAccessWhitelistResponseBody {
	return s.Body
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetStatusCode(v int32) *ModifyDBClusterAccessWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetBody(v *ModifyDBClusterAccessWhitelistResponseBody) *ModifyDBClusterAccessWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterAccessWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
