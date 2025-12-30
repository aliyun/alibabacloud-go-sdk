// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSyncEcsHostTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSyncEcsHostTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSyncEcsHostTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSyncEcsHostTaskResponseBody) *UpdateSyncEcsHostTaskResponse
	GetBody() *UpdateSyncEcsHostTaskResponseBody
}

type UpdateSyncEcsHostTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSyncEcsHostTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSyncEcsHostTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSyncEcsHostTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSyncEcsHostTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSyncEcsHostTaskResponse) GetBody() *UpdateSyncEcsHostTaskResponseBody {
	return s.Body
}

func (s *UpdateSyncEcsHostTaskResponse) SetHeaders(v map[string]*string) *UpdateSyncEcsHostTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateSyncEcsHostTaskResponse) SetStatusCode(v int32) *UpdateSyncEcsHostTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSyncEcsHostTaskResponse) SetBody(v *UpdateSyncEcsHostTaskResponseBody) *UpdateSyncEcsHostTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateSyncEcsHostTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
