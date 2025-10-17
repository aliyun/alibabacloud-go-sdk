// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterDeletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterDeletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterDeletionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterDeletionResponseBody) *ModifyDBClusterDeletionResponse
	GetBody() *ModifyDBClusterDeletionResponseBody
}

type ModifyDBClusterDeletionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterDeletionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterDeletionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDeletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterDeletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterDeletionResponse) GetBody() *ModifyDBClusterDeletionResponseBody {
	return s.Body
}

func (s *ModifyDBClusterDeletionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDeletionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDeletionResponse) SetStatusCode(v int32) *ModifyDBClusterDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterDeletionResponse) SetBody(v *ModifyDBClusterDeletionResponseBody) *ModifyDBClusterDeletionResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterDeletionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
