// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageDomainRoutingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStorageDomainRoutingRuleResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetStorageDomainRoutingRuleResponseBody
	GetCreateTime() *int64
	SetModifyTime(v int64) *GetStorageDomainRoutingRuleResponseBody
	GetModifyTime() *int64
	SetRequestId(v string) *GetStorageDomainRoutingRuleResponseBody
	GetRequestId() *string
	SetRoutes(v []*RouteItem) *GetStorageDomainRoutingRuleResponseBody
	GetRoutes() []*RouteItem
	SetRuleId(v string) *GetStorageDomainRoutingRuleResponseBody
	GetRuleId() *string
	SetSuccess(v bool) *GetStorageDomainRoutingRuleResponseBody
	GetSuccess() *bool
}

type GetStorageDomainRoutingRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1571926439000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE695******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing list.
	Routes []*RouteItem `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// crsdr-luq6qiegzvx****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStorageDomainRoutingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageDomainRoutingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetRoutes() []*RouteItem {
	return s.Routes
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *GetStorageDomainRoutingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetCode(v string) *GetStorageDomainRoutingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetCreateTime(v int64) *GetStorageDomainRoutingRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetModifyTime(v int64) *GetStorageDomainRoutingRuleResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetRequestId(v string) *GetStorageDomainRoutingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetRoutes(v []*RouteItem) *GetStorageDomainRoutingRuleResponseBody {
	s.Routes = v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetRuleId(v string) *GetStorageDomainRoutingRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) SetSuccess(v bool) *GetStorageDomainRoutingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetStorageDomainRoutingRuleResponseBody) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
