// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncThirdUserMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncThirdUserMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncThirdUserMappingResponse
	GetStatusCode() *int32
	SetBody(v *SyncThirdUserMappingResponseBody) *SyncThirdUserMappingResponse
	GetBody() *SyncThirdUserMappingResponseBody
}

type SyncThirdUserMappingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncThirdUserMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncThirdUserMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncThirdUserMappingResponse) GoString() string {
	return s.String()
}

func (s *SyncThirdUserMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncThirdUserMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncThirdUserMappingResponse) GetBody() *SyncThirdUserMappingResponseBody {
	return s.Body
}

func (s *SyncThirdUserMappingResponse) SetHeaders(v map[string]*string) *SyncThirdUserMappingResponse {
	s.Headers = v
	return s
}

func (s *SyncThirdUserMappingResponse) SetStatusCode(v int32) *SyncThirdUserMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncThirdUserMappingResponse) SetBody(v *SyncThirdUserMappingResponseBody) *SyncThirdUserMappingResponse {
	s.Body = v
	return s
}

func (s *SyncThirdUserMappingResponse) Validate() error {
	return dara.Validate(s)
}
