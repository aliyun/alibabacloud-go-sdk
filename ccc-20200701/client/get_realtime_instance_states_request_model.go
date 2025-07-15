// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealtimeInstanceStatesRequest
	GetInstanceId() *string
	SetMediaType(v string) *GetRealtimeInstanceStatesRequest
	GetMediaType() *string
}

type GetRealtimeInstanceStatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s GetRealtimeInstanceStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatesRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeInstanceStatesRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *GetRealtimeInstanceStatesRequest) SetInstanceId(v string) *GetRealtimeInstanceStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeInstanceStatesRequest) SetMediaType(v string) *GetRealtimeInstanceStatesRequest {
	s.MediaType = &v
	return s
}

func (s *GetRealtimeInstanceStatesRequest) Validate() error {
	return dara.Validate(s)
}
