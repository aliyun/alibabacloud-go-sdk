// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdInventoryConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*EpdInventoryConstituteItem) *GetEpdInventoryConstituteResponseBody
	GetData() []*EpdInventoryConstituteItem
	SetRequestId(v string) *GetEpdInventoryConstituteResponseBody
	GetRequestId() *string
}

type GetEpdInventoryConstituteResponseBody struct {
	// List of environmental impact results.
	Data []*EpdInventoryConstituteItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEpdInventoryConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEpdInventoryConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteResponseBody) GetData() []*EpdInventoryConstituteItem {
	return s.Data
}

func (s *GetEpdInventoryConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEpdInventoryConstituteResponseBody) SetData(v []*EpdInventoryConstituteItem) *GetEpdInventoryConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetEpdInventoryConstituteResponseBody) SetRequestId(v string) *GetEpdInventoryConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEpdInventoryConstituteResponseBody) Validate() error {
	return dara.Validate(s)
}
