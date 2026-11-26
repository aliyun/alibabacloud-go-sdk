// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUndeployHttpApiRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *BatchUndeployHttpApiRoutesRequest
	GetAsync() *bool
	SetRouteIds(v []*string) *BatchUndeployHttpApiRoutesRequest
	GetRouteIds() []*string
	SetUndeployAll(v bool) *BatchUndeployHttpApiRoutesRequest
	GetUndeployAll() *bool
}

type BatchUndeployHttpApiRoutesRequest struct {
	Async       *bool     `json:"async,omitempty" xml:"async,omitempty"`
	RouteIds    []*string `json:"routeIds,omitempty" xml:"routeIds,omitempty" type:"Repeated"`
	UndeployAll *bool     `json:"undeployAll,omitempty" xml:"undeployAll,omitempty"`
}

func (s BatchUndeployHttpApiRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUndeployHttpApiRoutesRequest) GoString() string {
	return s.String()
}

func (s *BatchUndeployHttpApiRoutesRequest) GetAsync() *bool {
	return s.Async
}

func (s *BatchUndeployHttpApiRoutesRequest) GetRouteIds() []*string {
	return s.RouteIds
}

func (s *BatchUndeployHttpApiRoutesRequest) GetUndeployAll() *bool {
	return s.UndeployAll
}

func (s *BatchUndeployHttpApiRoutesRequest) SetAsync(v bool) *BatchUndeployHttpApiRoutesRequest {
	s.Async = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesRequest) SetRouteIds(v []*string) *BatchUndeployHttpApiRoutesRequest {
	s.RouteIds = v
	return s
}

func (s *BatchUndeployHttpApiRoutesRequest) SetUndeployAll(v bool) *BatchUndeployHttpApiRoutesRequest {
	s.UndeployAll = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesRequest) Validate() error {
	return dara.Validate(s)
}
