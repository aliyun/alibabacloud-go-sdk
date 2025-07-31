// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckForOss(v bool) *QueryCallbackRequest
	GetCheckForOss() *bool
	SetId(v int64) *QueryCallbackRequest
	GetId() *int64
	SetRegionId(v string) *QueryCallbackRequest
	GetRegionId() *string
}

type QueryCallbackRequest struct {
	// example:
	//
	// true
	CheckForOss *bool `json:"CheckForOss,omitempty" xml:"CheckForOss,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackRequest) GoString() string {
	return s.String()
}

func (s *QueryCallbackRequest) GetCheckForOss() *bool {
	return s.CheckForOss
}

func (s *QueryCallbackRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryCallbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCallbackRequest) SetCheckForOss(v bool) *QueryCallbackRequest {
	s.CheckForOss = &v
	return s
}

func (s *QueryCallbackRequest) SetId(v int64) *QueryCallbackRequest {
	s.Id = &v
	return s
}

func (s *QueryCallbackRequest) SetRegionId(v string) *QueryCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCallbackRequest) Validate() error {
	return dara.Validate(s)
}
