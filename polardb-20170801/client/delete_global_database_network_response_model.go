// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalDatabaseNetworkResponseBody) *DeleteGlobalDatabaseNetworkResponse
	GetBody() *DeleteGlobalDatabaseNetworkResponseBody
}

type DeleteGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalDatabaseNetworkResponse) GetBody() *DeleteGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *DeleteGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *DeleteGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *DeleteGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkResponse) SetBody(v *DeleteGlobalDatabaseNetworkResponseBody) *DeleteGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalDatabaseNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
