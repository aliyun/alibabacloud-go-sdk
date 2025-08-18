// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*string) *GetEdgeContainerLogsResponseBody
	GetItems() []*string
	SetRequestId(v string) *GetEdgeContainerLogsResponseBody
	GetRequestId() *string
}

type GetEdgeContainerLogsResponseBody struct {
	// The logs.
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerLogsResponseBody) GetItems() []*string {
	return s.Items
}

func (s *GetEdgeContainerLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerLogsResponseBody) SetItems(v []*string) *GetEdgeContainerLogsResponseBody {
	s.Items = v
	return s
}

func (s *GetEdgeContainerLogsResponseBody) SetRequestId(v string) *GetEdgeContainerLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
