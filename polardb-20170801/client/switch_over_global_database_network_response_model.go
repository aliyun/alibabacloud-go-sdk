// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchOverGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchOverGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *SwitchOverGlobalDatabaseNetworkResponseBody) *SwitchOverGlobalDatabaseNetworkResponse
	GetBody() *SwitchOverGlobalDatabaseNetworkResponseBody
}

type SwitchOverGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchOverGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchOverGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) GetBody() *SwitchOverGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *SwitchOverGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *SwitchOverGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) SetBody(v *SwitchOverGlobalDatabaseNetworkResponseBody) *SwitchOverGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkResponse) Validate() error {
	return dara.Validate(s)
}
