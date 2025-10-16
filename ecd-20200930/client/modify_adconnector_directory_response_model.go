// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADConnectorDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyADConnectorDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyADConnectorDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyADConnectorDirectoryResponseBody) *ModifyADConnectorDirectoryResponse
	GetBody() *ModifyADConnectorDirectoryResponseBody
}

type ModifyADConnectorDirectoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyADConnectorDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyADConnectorDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyADConnectorDirectoryResponse) GetBody() *ModifyADConnectorDirectoryResponseBody {
	return s.Body
}

func (s *ModifyADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *ModifyADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorDirectoryResponse) SetStatusCode(v int32) *ModifyADConnectorDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyADConnectorDirectoryResponse) SetBody(v *ModifyADConnectorDirectoryResponseBody) *ModifyADConnectorDirectoryResponse {
	s.Body = v
	return s
}

func (s *ModifyADConnectorDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
