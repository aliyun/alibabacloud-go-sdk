// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerDeployRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*string) *GetEdgeContainerDeployRegionsResponseBody
	GetRegions() []*string
	SetRequestId(v string) *GetEdgeContainerDeployRegionsResponseBody
	GetRequestId() *string
}

type GetEdgeContainerDeployRegionsResponseBody struct {
	// The regions.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerDeployRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerDeployRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerDeployRegionsResponseBody) GetRegions() []*string {
	return s.Regions
}

func (s *GetEdgeContainerDeployRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerDeployRegionsResponseBody) SetRegions(v []*string) *GetEdgeContainerDeployRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *GetEdgeContainerDeployRegionsResponseBody) SetRequestId(v string) *GetEdgeContainerDeployRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerDeployRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
