// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSourceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupSourceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupSourceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupSourceEndpointResponseBody) *ModifyBackupSourceEndpointResponse
	GetBody() *ModifyBackupSourceEndpointResponseBody
}

type ModifyBackupSourceEndpointResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupSourceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupSourceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSourceEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupSourceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupSourceEndpointResponse) GetBody() *ModifyBackupSourceEndpointResponseBody {
	return s.Body
}

func (s *ModifyBackupSourceEndpointResponse) SetHeaders(v map[string]*string) *ModifyBackupSourceEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupSourceEndpointResponse) SetStatusCode(v int32) *ModifyBackupSourceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponse) SetBody(v *ModifyBackupSourceEndpointResponseBody) *ModifyBackupSourceEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupSourceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
