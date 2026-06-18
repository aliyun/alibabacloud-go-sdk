// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAssignSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *BatchAssignSeatsRequest
	GetAccountIds() []*string
	SetAccountIdsStr(v string) *BatchAssignSeatsRequest
	GetAccountIdsStr() *string
	SetCallerUacAccountId(v string) *BatchAssignSeatsRequest
	GetCallerUacAccountId() *string
	SetLocale(v string) *BatchAssignSeatsRequest
	GetLocale() *string
	SetNamespaceId(v string) *BatchAssignSeatsRequest
	GetNamespaceId() *string
	SetSeatType(v string) *BatchAssignSeatsRequest
	GetSeatType() *string
	SetWorkspaceId(v string) *BatchAssignSeatsRequest
	GetWorkspaceId() *string
}

type BatchAssignSeatsRequest struct {
	// The list of target member IDs.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The AccountIds in string format.
	//
	// example:
	//
	// [\\"5854302538759072\\"]
	AccountIdsStr *string `json:"AccountIdsStr,omitempty" xml:"AccountIdsStr,omitempty"`
	// The account ID of the caller that identifies the initiator of this call.
	//
	// example:
	//
	// acc_123456789
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// The language setting. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The product namespace ID. For the TokenPlan product, this field is fixed to namespace-1.
	//
	// example:
	//
	// namespace-1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The seat type. Valid values: standard, pro, and max.
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	SeatType *string `json:"SeatType,omitempty" xml:"SeatType,omitempty"`
	// The ID of the target workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws_123456789
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BatchAssignSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAssignSeatsRequest) GoString() string {
	return s.String()
}

func (s *BatchAssignSeatsRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *BatchAssignSeatsRequest) GetAccountIdsStr() *string {
	return s.AccountIdsStr
}

func (s *BatchAssignSeatsRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *BatchAssignSeatsRequest) GetLocale() *string {
	return s.Locale
}

func (s *BatchAssignSeatsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *BatchAssignSeatsRequest) GetSeatType() *string {
	return s.SeatType
}

func (s *BatchAssignSeatsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchAssignSeatsRequest) SetAccountIds(v []*string) *BatchAssignSeatsRequest {
	s.AccountIds = v
	return s
}

func (s *BatchAssignSeatsRequest) SetAccountIdsStr(v string) *BatchAssignSeatsRequest {
	s.AccountIdsStr = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetCallerUacAccountId(v string) *BatchAssignSeatsRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetLocale(v string) *BatchAssignSeatsRequest {
	s.Locale = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetNamespaceId(v string) *BatchAssignSeatsRequest {
	s.NamespaceId = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetSeatType(v string) *BatchAssignSeatsRequest {
	s.SeatType = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetWorkspaceId(v string) *BatchAssignSeatsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchAssignSeatsRequest) Validate() error {
	return dara.Validate(s)
}
