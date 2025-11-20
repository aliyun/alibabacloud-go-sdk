// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryConferenceInfoByRoomCodeRequestTenantContext) *QueryConferenceInfoByRoomCodeRequest
	GetTenantContext() *QueryConferenceInfoByRoomCodeRequestTenantContext
	SetMaxResults(v int32) *QueryConferenceInfoByRoomCodeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryConferenceInfoByRoomCodeRequest
	GetNextToken() *string
	SetRoomCode(v string) *QueryConferenceInfoByRoomCodeRequest
	GetRoomCode() *string
}

type QueryConferenceInfoByRoomCodeRequest struct {
	TenantContext *QueryConferenceInfoByRoomCodeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 583 480 813
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeRequest) GetTenantContext() *QueryConferenceInfoByRoomCodeRequestTenantContext {
	return s.TenantContext
}

func (s *QueryConferenceInfoByRoomCodeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryConferenceInfoByRoomCodeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceInfoByRoomCodeRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetTenantContext(v *QueryConferenceInfoByRoomCodeRequestTenantContext) *QueryConferenceInfoByRoomCodeRequest {
	s.TenantContext = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetMaxResults(v int32) *QueryConferenceInfoByRoomCodeRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetNextToken(v string) *QueryConferenceInfoByRoomCodeRequest {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetRoomCode(v string) *QueryConferenceInfoByRoomCodeRequest {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConferenceInfoByRoomCodeRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryConferenceInfoByRoomCodeRequestTenantContext) SetTenantId(v string) *QueryConferenceInfoByRoomCodeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
