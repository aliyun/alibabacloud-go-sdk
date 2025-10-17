// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalDatabaseNetworkResponseBody) *CreateGlobalDatabaseNetworkResponse
	GetBody() *CreateGlobalDatabaseNetworkResponseBody
}

type CreateGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalDatabaseNetworkResponse) GetBody() *CreateGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *CreateGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *CreateGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *CreateGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponse) SetBody(v *CreateGlobalDatabaseNetworkResponseBody) *CreateGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
