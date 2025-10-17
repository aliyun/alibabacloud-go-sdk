// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *ResetGlobalDatabaseNetworkResponseBody) *ResetGlobalDatabaseNetworkResponse
	GetBody() *ResetGlobalDatabaseNetworkResponseBody
}

type ResetGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *ResetGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetGlobalDatabaseNetworkResponse) GetBody() *ResetGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *ResetGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *ResetGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *ResetGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *ResetGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkResponse) SetBody(v *ResetGlobalDatabaseNetworkResponseBody) *ResetGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *ResetGlobalDatabaseNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
