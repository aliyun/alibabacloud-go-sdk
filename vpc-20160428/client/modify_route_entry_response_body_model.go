// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouteEntryResponseBody
	GetRequestId() *string
}

type ModifyRouteEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 861E6630-AEC0-4B2D-B214-6CB5E44B7F04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouteEntryResponseBody) SetRequestId(v string) *ModifyRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
