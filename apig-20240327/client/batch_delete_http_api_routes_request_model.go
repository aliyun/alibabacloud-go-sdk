// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHttpApiRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAll(v bool) *BatchDeleteHttpApiRoutesRequest
	GetDeleteAll() *bool
	SetRouteIds(v []*string) *BatchDeleteHttpApiRoutesRequest
	GetRouteIds() []*string
}

type BatchDeleteHttpApiRoutesRequest struct {
	DeleteAll *bool     `json:"deleteAll,omitempty" xml:"deleteAll,omitempty"`
	RouteIds  []*string `json:"routeIds,omitempty" xml:"routeIds,omitempty" type:"Repeated"`
}

func (s BatchDeleteHttpApiRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHttpApiRoutesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteHttpApiRoutesRequest) GetDeleteAll() *bool {
	return s.DeleteAll
}

func (s *BatchDeleteHttpApiRoutesRequest) GetRouteIds() []*string {
	return s.RouteIds
}

func (s *BatchDeleteHttpApiRoutesRequest) SetDeleteAll(v bool) *BatchDeleteHttpApiRoutesRequest {
	s.DeleteAll = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesRequest) SetRouteIds(v []*string) *BatchDeleteHttpApiRoutesRequest {
	s.RouteIds = v
	return s
}

func (s *BatchDeleteHttpApiRoutesRequest) Validate() error {
	return dara.Validate(s)
}
