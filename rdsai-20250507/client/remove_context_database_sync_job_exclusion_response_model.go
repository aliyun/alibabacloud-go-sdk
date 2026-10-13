// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContextDatabaseSyncJobExclusionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveContextDatabaseSyncJobExclusionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveContextDatabaseSyncJobExclusionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveContextDatabaseSyncJobExclusionResponseBody) *RemoveContextDatabaseSyncJobExclusionResponse
	GetBody() *RemoveContextDatabaseSyncJobExclusionResponseBody
}

type RemoveContextDatabaseSyncJobExclusionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveContextDatabaseSyncJobExclusionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveContextDatabaseSyncJobExclusionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveContextDatabaseSyncJobExclusionResponse) GoString() string {
	return s.String()
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) GetBody() *RemoveContextDatabaseSyncJobExclusionResponseBody {
	return s.Body
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) SetHeaders(v map[string]*string) *RemoveContextDatabaseSyncJobExclusionResponse {
	s.Headers = v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) SetStatusCode(v int32) *RemoveContextDatabaseSyncJobExclusionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) SetBody(v *RemoveContextDatabaseSyncJobExclusionResponseBody) *RemoveContextDatabaseSyncJobExclusionResponse {
	s.Body = v
	return s
}

func (s *RemoveContextDatabaseSyncJobExclusionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
