// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsoleMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DescribeConsoleMessagesRequest
	GetBusinessChannel() *string
	SetClientTypes(v []*string) *DescribeConsoleMessagesRequest
	GetClientTypes() []*string
	SetKeyword(v string) *DescribeConsoleMessagesRequest
	GetKeyword() *string
	SetMessageIds(v []*string) *DescribeConsoleMessagesRequest
	GetMessageIds() []*string
	SetMessageLevel(v string) *DescribeConsoleMessagesRequest
	GetMessageLevel() *string
	SetMessageType(v string) *DescribeConsoleMessagesRequest
	GetMessageType() *string
	SetPageNum(v int32) *DescribeConsoleMessagesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeConsoleMessagesRequest
	GetPageSize() *int32
	SetTargetIds(v []*string) *DescribeConsoleMessagesRequest
	GetTargetIds() []*string
	SetTargetType(v string) *DescribeConsoleMessagesRequest
	GetTargetType() *string
}

type DescribeConsoleMessagesRequest struct {
	// The business channel, which is used to restrict the business to which the messages to query belong. Valid values: enterprise, standard, office, business, lincore_team, and AGENTIC_CLOUD_COMPUTER. The values are not case-sensitive. Default value: enterprise. Unrecognized channel values are also processed as enterprise.
	//
	// example:
	//
	// Enterprise
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The filter list of client types to which operational messages apply. A maximum of 100 items can be included. Valid values: Windows (Windows client), Mac (macOS client), Web (web client), and Hardware (hardware terminal). For Mac, macOS is also supported. The matching is not case-sensitive. If multiple types are specified, a match occurs if any type matches. Do not pass the same type repeatedly. To filter by client type, provide a non-empty list. If this parameter is not specified, messages are not filtered by client type. Do not specify this parameter when querying NOTIFICATION messages.
	ClientTypes []*string `json:"ClientTypes,omitempty" xml:"ClientTypes,omitempty" type:"Repeated"`
	// The search keyword, which is used to perform a fuzzy match on the message title or body. A match occurs if either the title or the body matches the keyword. If this parameter is not specified or an empty string is provided, messages are not filtered by keyword.
	//
	// example:
	//
	// System maintenance
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The list of message IDs to query. A maximum of 100 message IDs can be included. If multiple IDs are provided, a match occurs if any ID matches. If this parameter is not specified or an empty list is provided, messages are not filtered by message ID. If no message ID matches, an empty list is returned.
	MessageIds []*string `json:"MessageIds,omitempty" xml:"MessageIds,omitempty" type:"Repeated"`
	// The message level filter condition for notification messages. Valid values: STRONG (strong message) and WEAK (weak message). The values are not case-sensitive. If this parameter is not specified or the value is unrecognized, messages are not filtered by message level. Do not specify this parameter when querying OPERATIONAL messages.
	//
	// example:
	//
	// STRONG
	MessageLevel *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
	// The message type. Valid values: OPERATIONAL (client operational message), WALLPAPER (desktop wallpaper text message), NOTIFICATION (message center notification), and CUSTOM_WALLPAPER (custom desktop wallpaper). The values are not case-sensitive. If this parameter is not specified or the value is unrecognized, messages are not filtered by message type.
	//
	// example:
	//
	// NOTIFICATION
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The page number of the query result. Pages start from 1. Default value: 1. If the value is less than 1, it is processed as 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values: 1 to 200. Default value: 200. If the value is less than 1 or greater than 200, it is processed as 200.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of target IDs to match. A maximum of 100 items can be included. To filter by target, provide a non-empty list and specify TargetType at the same time. Currently, only the first ID is used for Organization, EndUser, and AdUser. Queries for EndUser and AdUser also match messages targeted at the organization to which the user belongs and all users. Queries for Organization include messages for all convenience users and are compatible with fuzzy matching of organization IDs. Other target matching methods match based on target IDs, and the target type of the returned records is not guaranteed to be consistent with TargetType. Currently, Tag does not perform target ID filtering.
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
	// The target matching method. Valid values: AllUsers (all convenience users), EndUser (specified convenience user), Organization (convenience user organization), Desktop (cloud computer), ResourceGroup (resource group), Region (region), Tag (tag), AdAllUsers (all AD users), AdUser (specified AD user), and AdOrg (AD organization). The values are not case-sensitive. If TargetIds is specified, this parameter must also be specified. The target type in the query results is not guaranteed to be consistent with this parameter. If only this parameter is specified without target IDs, no target filtering is applied. Currently, Tag does not perform target filtering.
	//
	// example:
	//
	// EndUser
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeConsoleMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeConsoleMessagesRequest) GetClientTypes() []*string {
	return s.ClientTypes
}

func (s *DescribeConsoleMessagesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeConsoleMessagesRequest) GetMessageIds() []*string {
	return s.MessageIds
}

func (s *DescribeConsoleMessagesRequest) GetMessageLevel() *string {
	return s.MessageLevel
}

func (s *DescribeConsoleMessagesRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *DescribeConsoleMessagesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeConsoleMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsoleMessagesRequest) GetTargetIds() []*string {
	return s.TargetIds
}

func (s *DescribeConsoleMessagesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeConsoleMessagesRequest) SetBusinessChannel(v string) *DescribeConsoleMessagesRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetClientTypes(v []*string) *DescribeConsoleMessagesRequest {
	s.ClientTypes = v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetKeyword(v string) *DescribeConsoleMessagesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetMessageIds(v []*string) *DescribeConsoleMessagesRequest {
	s.MessageIds = v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetMessageLevel(v string) *DescribeConsoleMessagesRequest {
	s.MessageLevel = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetMessageType(v string) *DescribeConsoleMessagesRequest {
	s.MessageType = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetPageNum(v int32) *DescribeConsoleMessagesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetPageSize(v int32) *DescribeConsoleMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetTargetIds(v []*string) *DescribeConsoleMessagesRequest {
	s.TargetIds = v
	return s
}

func (s *DescribeConsoleMessagesRequest) SetTargetType(v string) *DescribeConsoleMessagesRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeConsoleMessagesRequest) Validate() error {
	return dara.Validate(s)
}
