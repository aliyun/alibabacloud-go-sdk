// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelContextDatabaseSyncRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelContextDatabaseSyncRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelContextDatabaseSyncRunResponse
	GetStatusCode() *int32
	SetBody(v *CancelContextDatabaseSyncRunResponseBody) *CancelContextDatabaseSyncRunResponse
	GetBody() *CancelContextDatabaseSyncRunResponseBody
}

type CancelContextDatabaseSyncRunResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelContextDatabaseSyncRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelContextDatabaseSyncRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelContextDatabaseSyncRunResponse) GoString() string {
	return s.String()
}

func (s *CancelContextDatabaseSyncRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelContextDatabaseSyncRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelContextDatabaseSyncRunResponse) GetBody() *CancelContextDatabaseSyncRunResponseBody {
	return s.Body
}

func (s *CancelContextDatabaseSyncRunResponse) SetHeaders(v map[string]*string) *CancelContextDatabaseSyncRunResponse {
	s.Headers = v
	return s
}

func (s *CancelContextDatabaseSyncRunResponse) SetStatusCode(v int32) *CancelContextDatabaseSyncRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelContextDatabaseSyncRunResponse) SetBody(v *CancelContextDatabaseSyncRunResponseBody) *CancelContextDatabaseSyncRunResponse {
	s.Body = v
	return s
}

func (s *CancelContextDatabaseSyncRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
