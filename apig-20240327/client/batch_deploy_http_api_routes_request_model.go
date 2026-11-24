// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployHttpApiRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *BatchDeployHttpApiRoutesRequest
	GetAsync() *bool
	SetDeployAll(v bool) *BatchDeployHttpApiRoutesRequest
	GetDeployAll() *bool
	SetRouteIds(v []*string) *BatchDeployHttpApiRoutesRequest
	GetRouteIds() []*string
}

type BatchDeployHttpApiRoutesRequest struct {
	// Specifies whether to publish routes asynchronously. Default value: true. If set to true, the acceptance results of the routes are returned. If set to false, the system waits for the batch processing to complete and returns the final results and statistics for each route.
	//
	// example:
	//
	// false
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// Specifies whether to publish all routes visible to the current user under the API. Default value: false. If you set this parameter to true, all visible routes are published and the routeIds parameter is ignored.
	//
	// example:
	//
	// false
	DeployAll *bool `json:"deployAll,omitempty" xml:"deployAll,omitempty"`
	// The list of route IDs to publish. If deployAll is omitted or set to false, at least one route ID must be provided. If any route does not exist or belongs to another API, the entire batch request is rejected.
	//
	// example:
	//
	// ["hr-example001","hr-example002"]
	RouteIds []*string `json:"routeIds,omitempty" xml:"routeIds,omitempty" type:"Repeated"`
}

func (s BatchDeployHttpApiRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployHttpApiRoutesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeployHttpApiRoutesRequest) GetAsync() *bool {
	return s.Async
}

func (s *BatchDeployHttpApiRoutesRequest) GetDeployAll() *bool {
	return s.DeployAll
}

func (s *BatchDeployHttpApiRoutesRequest) GetRouteIds() []*string {
	return s.RouteIds
}

func (s *BatchDeployHttpApiRoutesRequest) SetAsync(v bool) *BatchDeployHttpApiRoutesRequest {
	s.Async = &v
	return s
}

func (s *BatchDeployHttpApiRoutesRequest) SetDeployAll(v bool) *BatchDeployHttpApiRoutesRequest {
	s.DeployAll = &v
	return s
}

func (s *BatchDeployHttpApiRoutesRequest) SetRouteIds(v []*string) *BatchDeployHttpApiRoutesRequest {
	s.RouteIds = v
	return s
}

func (s *BatchDeployHttpApiRoutesRequest) Validate() error {
	return dara.Validate(s)
}
