// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceGroupForServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenServiceGroupForServiceResponseBody
	GetRequestId() *string
	SetServiceIds(v []*int64) *OpenServiceGroupForServiceResponseBody
	GetServiceIds() []*int64
}

type OpenServiceGroupForServiceResponseBody struct {
	// example:
	//
	// 3239281273464326823
	RequestId  *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceIds []*int64 `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s OpenServiceGroupForServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceGroupForServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceGroupForServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenServiceGroupForServiceResponseBody) GetServiceIds() []*int64 {
	return s.ServiceIds
}

func (s *OpenServiceGroupForServiceResponseBody) SetRequestId(v string) *OpenServiceGroupForServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenServiceGroupForServiceResponseBody) SetServiceIds(v []*int64) *OpenServiceGroupForServiceResponseBody {
	s.ServiceIds = v
	return s
}

func (s *OpenServiceGroupForServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
