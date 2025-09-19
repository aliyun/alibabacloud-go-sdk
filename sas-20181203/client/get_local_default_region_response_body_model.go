// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalDefaultRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetLocalDefaultRegionResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetLocalDefaultRegionResponseBody
	GetStatus() *int32
}

type GetLocalDefaultRegionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6D23860-55C2-55AA-B484-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The access type of the multi-cloud site. Valid values:
	//
	// 	- **0**: The current site is not the default site of the multi-cloud site. You can specify a site as the default site of the multi-cloud site.
	//
	// 	- **1**: The current site is the default site of the multi-cloud site.
	//
	// 	- **2**: Another site is set as the default site of the multi-cloud site.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLocalDefaultRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLocalDefaultRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetLocalDefaultRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLocalDefaultRegionResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetLocalDefaultRegionResponseBody) SetRequestId(v string) *GetLocalDefaultRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLocalDefaultRegionResponseBody) SetStatus(v int32) *GetLocalDefaultRegionResponseBody {
	s.Status = &v
	return s
}

func (s *GetLocalDefaultRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
