// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeRouterWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeRouterWeightResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyDBNodeRouterWeightResponseBody
	GetRequestId() *string
}

type ModifyDBNodeRouterWeightResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E2FDB684-751D-424D-98B9-704BEA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeRouterWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeRouterWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeRouterWeightResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeRouterWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeRouterWeightResponseBody) SetDBClusterId(v string) *ModifyDBNodeRouterWeightResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeRouterWeightResponseBody) SetRequestId(v string) *ModifyDBNodeRouterWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeRouterWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
