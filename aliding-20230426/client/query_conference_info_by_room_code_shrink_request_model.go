// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryConferenceInfoByRoomCodeShrinkRequest
	GetTenantContextShrink() *string
	SetMaxResults(v int32) *QueryConferenceInfoByRoomCodeShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryConferenceInfoByRoomCodeShrinkRequest
	GetNextToken() *string
	SetRoomCode(v string) *QueryConferenceInfoByRoomCodeShrinkRequest
	GetRoomCode() *string
}

type QueryConferenceInfoByRoomCodeShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
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

func (s QueryConferenceInfoByRoomCodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) SetTenantContextShrink(v string) *QueryConferenceInfoByRoomCodeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) SetMaxResults(v int32) *QueryConferenceInfoByRoomCodeShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) SetNextToken(v string) *QueryConferenceInfoByRoomCodeShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) SetRoomCode(v string) *QueryConferenceInfoByRoomCodeShrinkRequest {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
