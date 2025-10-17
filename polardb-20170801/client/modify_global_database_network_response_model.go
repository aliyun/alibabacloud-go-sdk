// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalDatabaseNetworkResponseBody) *ModifyGlobalDatabaseNetworkResponse
	GetBody() *ModifyGlobalDatabaseNetworkResponseBody
}

type ModifyGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalDatabaseNetworkResponse) GetBody() *ModifyGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *ModifyGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *ModifyGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *ModifyGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkResponse) SetBody(v *ModifyGlobalDatabaseNetworkResponseBody) *ModifyGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalDatabaseNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
