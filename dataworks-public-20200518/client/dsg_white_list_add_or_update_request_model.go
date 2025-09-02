// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListAddOrUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWhiteLists(v []*DsgWhiteListAddOrUpdateRequestWhiteLists) *DsgWhiteListAddOrUpdateRequest
	GetWhiteLists() []*DsgWhiteListAddOrUpdateRequestWhiteLists
}

type DsgWhiteListAddOrUpdateRequest struct {
	// A collection of whitelists.
	//
	// This parameter is required.
	WhiteLists []*DsgWhiteListAddOrUpdateRequestWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s DsgWhiteListAddOrUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListAddOrUpdateRequest) GoString() string {
	return s.String()
}

func (s *DsgWhiteListAddOrUpdateRequest) GetWhiteLists() []*DsgWhiteListAddOrUpdateRequestWhiteLists {
	return s.WhiteLists
}

func (s *DsgWhiteListAddOrUpdateRequest) SetWhiteLists(v []*DsgWhiteListAddOrUpdateRequestWhiteLists) *DsgWhiteListAddOrUpdateRequest {
	s.WhiteLists = v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequest) Validate() error {
	return dara.Validate(s)
}

type DsgWhiteListAddOrUpdateRequestWhiteLists struct {
	// The end of the time range to query. If you enter null, the whitelist is valid permanently.
	//
	// example:
	//
	// null
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the data masking whitelist.
	//
	// 	- If you do not configure this parameter, the current operation is to add a data masking whitelist.
	//
	// 	- If you configure this parameter, the current operation is to modify a data masking whitelist. You can call the [DsgWhiteListQueryList](https://help.aliyun.com/document_detail/2786508.html) operation to query the whitelist ID.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the data masking rule. You can call the [DsgDesensPlanQueryList](https://help.aliyun.com/document_detail/2786578.html) operation to query the ID of the data masking rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-04-10 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A collection of user group IDs.
	//
	// This parameter is required.
	UserGroupIds []*int32 `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s DsgWhiteListAddOrUpdateRequestWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListAddOrUpdateRequestWhiteLists) GoString() string {
	return s.String()
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) GetEndTime() *string {
	return s.EndTime
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) GetId() *int32 {
	return s.Id
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) GetRuleId() *int32 {
	return s.RuleId
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) GetStartTime() *string {
	return s.StartTime
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) GetUserGroupIds() []*int32 {
	return s.UserGroupIds
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) SetEndTime(v string) *DsgWhiteListAddOrUpdateRequestWhiteLists {
	s.EndTime = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) SetId(v int32) *DsgWhiteListAddOrUpdateRequestWhiteLists {
	s.Id = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) SetRuleId(v int32) *DsgWhiteListAddOrUpdateRequestWhiteLists {
	s.RuleId = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) SetStartTime(v string) *DsgWhiteListAddOrUpdateRequestWhiteLists {
	s.StartTime = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) SetUserGroupIds(v []*int32) *DsgWhiteListAddOrUpdateRequestWhiteLists {
	s.UserGroupIds = v
	return s
}

func (s *DsgWhiteListAddOrUpdateRequestWhiteLists) Validate() error {
	return dara.Validate(s)
}
