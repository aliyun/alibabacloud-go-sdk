// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFigureClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFigureCluster(v *FigureCluster) *GetFigureClusterResponseBody
	GetFigureCluster() *FigureCluster
	SetRequestId(v string) *GetFigureClusterResponseBody
	GetRequestId() *string
}

type GetFigureClusterResponseBody struct {
	// The information about the face cluster.
	FigureCluster *FigureCluster `json:"FigureCluster,omitempty" xml:"FigureCluster,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFigureClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetFigureClusterResponseBody) GetFigureCluster() *FigureCluster {
	return s.FigureCluster
}

func (s *GetFigureClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFigureClusterResponseBody) SetFigureCluster(v *FigureCluster) *GetFigureClusterResponseBody {
	s.FigureCluster = v
	return s
}

func (s *GetFigureClusterResponseBody) SetRequestId(v string) *GetFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFigureClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
