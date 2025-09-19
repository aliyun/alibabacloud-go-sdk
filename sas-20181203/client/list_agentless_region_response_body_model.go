// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionList(v []*string) *ListAgentlessRegionResponseBody
	GetRegionList() []*string
	SetRequestId(v string) *ListAgentlessRegionResponseBody
	GetRequestId() *string
}

type ListAgentlessRegionResponseBody struct {
	// The information about the regions.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessRegionResponseBody) GetRegionList() []*string {
	return s.RegionList
}

func (s *ListAgentlessRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessRegionResponseBody) SetRegionList(v []*string) *ListAgentlessRegionResponseBody {
	s.RegionList = v
	return s
}

func (s *ListAgentlessRegionResponseBody) SetRequestId(v string) *ListAgentlessRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
