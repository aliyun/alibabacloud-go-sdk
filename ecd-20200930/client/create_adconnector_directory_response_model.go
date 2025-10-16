// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateADConnectorDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateADConnectorDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateADConnectorDirectoryResponseBody) *CreateADConnectorDirectoryResponse
	GetBody() *CreateADConnectorDirectoryResponseBody
}

type CreateADConnectorDirectoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateADConnectorDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateADConnectorDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateADConnectorDirectoryResponse) GetBody() *CreateADConnectorDirectoryResponseBody {
	return s.Body
}

func (s *CreateADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *CreateADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetStatusCode(v int32) *CreateADConnectorDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetBody(v *CreateADConnectorDirectoryResponseBody) *CreateADConnectorDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateADConnectorDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
