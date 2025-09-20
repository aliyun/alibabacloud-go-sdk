// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFigureClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFigureClusters(v []*FigureCluster) *BatchGetFigureClusterResponseBody
	GetFigureClusters() []*FigureCluster
	SetRequestId(v string) *BatchGetFigureClusterResponseBody
	GetRequestId() *string
}

type BatchGetFigureClusterResponseBody struct {
	// The clusters.
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetFigureClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFigureClusterResponseBody) GetFigureClusters() []*FigureCluster {
	return s.FigureClusters
}

func (s *BatchGetFigureClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetFigureClusterResponseBody) SetFigureClusters(v []*FigureCluster) *BatchGetFigureClusterResponseBody {
	s.FigureClusters = v
	return s
}

func (s *BatchGetFigureClusterResponseBody) SetRequestId(v string) *BatchGetFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetFigureClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
