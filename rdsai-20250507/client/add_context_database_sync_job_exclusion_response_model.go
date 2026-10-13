// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContextDatabaseSyncJobExclusionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContextDatabaseSyncJobExclusionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContextDatabaseSyncJobExclusionResponse
	GetStatusCode() *int32
	SetBody(v *AddContextDatabaseSyncJobExclusionResponseBody) *AddContextDatabaseSyncJobExclusionResponse
	GetBody() *AddContextDatabaseSyncJobExclusionResponseBody
}

type AddContextDatabaseSyncJobExclusionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContextDatabaseSyncJobExclusionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContextDatabaseSyncJobExclusionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContextDatabaseSyncJobExclusionResponse) GoString() string {
	return s.String()
}

func (s *AddContextDatabaseSyncJobExclusionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContextDatabaseSyncJobExclusionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContextDatabaseSyncJobExclusionResponse) GetBody() *AddContextDatabaseSyncJobExclusionResponseBody {
	return s.Body
}

func (s *AddContextDatabaseSyncJobExclusionResponse) SetHeaders(v map[string]*string) *AddContextDatabaseSyncJobExclusionResponse {
	s.Headers = v
	return s
}

func (s *AddContextDatabaseSyncJobExclusionResponse) SetStatusCode(v int32) *AddContextDatabaseSyncJobExclusionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContextDatabaseSyncJobExclusionResponse) SetBody(v *AddContextDatabaseSyncJobExclusionResponseBody) *AddContextDatabaseSyncJobExclusionResponse {
	s.Body = v
	return s
}

func (s *AddContextDatabaseSyncJobExclusionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
