// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncOrgStructureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncOrgStructureResponse
	GetStatusCode() *int32
	SetBody(v *SyncOrgStructureResponseBody) *SyncOrgStructureResponse
	GetBody() *SyncOrgStructureResponseBody
}

type SyncOrgStructureResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncOrgStructureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncOrgStructureResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureResponse) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncOrgStructureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncOrgStructureResponse) GetBody() *SyncOrgStructureResponseBody {
	return s.Body
}

func (s *SyncOrgStructureResponse) SetHeaders(v map[string]*string) *SyncOrgStructureResponse {
	s.Headers = v
	return s
}

func (s *SyncOrgStructureResponse) SetStatusCode(v int32) *SyncOrgStructureResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncOrgStructureResponse) SetBody(v *SyncOrgStructureResponseBody) *SyncOrgStructureResponse {
	s.Body = v
	return s
}

func (s *SyncOrgStructureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
