// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDeletionProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterDeletionProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterDeletionProtectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterDeletionProtectionResponseBody) *ModifyClusterDeletionProtectionResponse
	GetBody() *ModifyClusterDeletionProtectionResponseBody
}

type ModifyClusterDeletionProtectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDeletionProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterDeletionProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterDeletionProtectionResponse) GetBody() *ModifyClusterDeletionProtectionResponseBody {
	return s.Body
}

func (s *ModifyClusterDeletionProtectionResponse) SetHeaders(v map[string]*string) *ModifyClusterDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) SetStatusCode(v int32) *ModifyClusterDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) SetBody(v *ModifyClusterDeletionProtectionResponseBody) *ModifyClusterDeletionProtectionResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
