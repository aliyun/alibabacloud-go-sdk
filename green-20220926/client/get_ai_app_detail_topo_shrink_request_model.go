// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailTopoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailTopoShrinkRequest
	GetAppId() *string
	SetRegionId(v string) *GetAiAppDetailTopoShrinkRequest
	GetRegionId() *string
	SetTimeQueryShrink(v string) *GetAiAppDetailTopoShrinkRequest
	GetTimeQueryShrink() *string
}

type GetAiAppDetailTopoShrinkRequest struct {
	// The application ID that identifies a specific AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time query.
	TimeQueryShrink *string `json:"TimeQuery,omitempty" xml:"TimeQuery,omitempty"`
}

func (s GetAiAppDetailTopoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailTopoShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppDetailTopoShrinkRequest) GetTimeQueryShrink() *string {
	return s.TimeQueryShrink
}

func (s *GetAiAppDetailTopoShrinkRequest) SetAppId(v string) *GetAiAppDetailTopoShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailTopoShrinkRequest) SetRegionId(v string) *GetAiAppDetailTopoShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppDetailTopoShrinkRequest) SetTimeQueryShrink(v string) *GetAiAppDetailTopoShrinkRequest {
	s.TimeQueryShrink = &v
	return s
}

func (s *GetAiAppDetailTopoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
