// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentAlertRulesResponseBody
	GetCode() *int32
	SetData(v *ListEnvironmentAlertRulesResponseBodyData) *ListEnvironmentAlertRulesResponseBody
	GetData() *ListEnvironmentAlertRulesResponseBodyData
	SetMessage(v string) *ListEnvironmentAlertRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentAlertRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentAlertRulesResponseBody
	GetSuccess() *bool
}

type ListEnvironmentAlertRulesResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *ListEnvironmentAlertRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAlertRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentAlertRulesResponseBody) GetData() *ListEnvironmentAlertRulesResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentAlertRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentAlertRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentAlertRulesResponseBody) SetCode(v int32) *ListEnvironmentAlertRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBody) SetData(v *ListEnvironmentAlertRulesResponseBodyData) *ListEnvironmentAlertRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBody) SetMessage(v string) *ListEnvironmentAlertRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBody) SetRequestId(v string) *ListEnvironmentAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBody) SetSuccess(v bool) *ListEnvironmentAlertRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEnvironmentAlertRulesResponseBodyData struct {
	// The queried alert groups.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The queried rules.
	Rules []*ListEnvironmentAlertRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 26
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEnvironmentAlertRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAlertRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAlertRulesResponseBodyData) GetGroups() []*string {
	return s.Groups
}

func (s *ListEnvironmentAlertRulesResponseBodyData) GetRules() []*ListEnvironmentAlertRulesResponseBodyDataRules {
	return s.Rules
}

func (s *ListEnvironmentAlertRulesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListEnvironmentAlertRulesResponseBodyData) SetGroups(v []*string) *ListEnvironmentAlertRulesResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBodyData) SetRules(v []*ListEnvironmentAlertRulesResponseBodyDataRules) *ListEnvironmentAlertRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBodyData) SetTotal(v int64) *ListEnvironmentAlertRulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBodyData) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvironmentAlertRulesResponseBodyDataRules struct {
	// The ID of the alert rule.
	//
	// example:
	//
	// 9502571
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// mysql-CS-MySQLInnoDBLogWaits_lu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEnvironmentAlertRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAlertRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAlertRulesResponseBodyDataRules) GetAlertId() *int64 {
	return s.AlertId
}

func (s *ListEnvironmentAlertRulesResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *ListEnvironmentAlertRulesResponseBodyDataRules) SetAlertId(v int64) *ListEnvironmentAlertRulesResponseBodyDataRules {
	s.AlertId = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBodyDataRules) SetName(v string) *ListEnvironmentAlertRulesResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *ListEnvironmentAlertRulesResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
