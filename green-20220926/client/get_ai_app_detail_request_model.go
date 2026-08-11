// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailRequest
	GetAppId() *string
	SetRegionId(v string) *GetAiAppDetailRequest
	GetRegionId() *string
}

type GetAiAppDetailRequest struct {
	// The ID of the AI application. This parameter is required.
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
}

func (s GetAiAppDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppDetailRequest) SetAppId(v string) *GetAiAppDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailRequest) SetRegionId(v string) *GetAiAppDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppDetailRequest) Validate() error {
	return dara.Validate(s)
}
