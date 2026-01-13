// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountConfigInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountStatus(v string) *GetAccountConfigInfoResponseBody
	GetAccountStatus() *string
	SetAvailableTime(v string) *GetAccountConfigInfoResponseBody
	GetAvailableTime() *string
	SetBillQps(v string) *GetAccountConfigInfoResponseBody
	GetBillQps() *string
	SetLadderType(v string) *GetAccountConfigInfoResponseBody
	GetLadderType() *string
	SetMainAccountId(v string) *GetAccountConfigInfoResponseBody
	GetMainAccountId() *string
	SetRequestId(v string) *GetAccountConfigInfoResponseBody
	GetRequestId() *string
}

type GetAccountConfigInfoResponseBody struct {
	// example:
	//
	// Normal
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// example:
	//
	// 2025-07-12 12:00:00
	AvailableTime *string `json:"availableTime,omitempty" xml:"availableTime,omitempty"`
	// example:
	//
	// 30
	BillQps *string `json:"billQps,omitempty" xml:"billQps,omitempty"`
	// example:
	//
	// FixLadder
	LadderType *string `json:"ladderType,omitempty" xml:"ladderType,omitempty"`
	// example:
	//
	// 12123123123
	MainAccountId *string `json:"mainAccountId,omitempty" xml:"mainAccountId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6788a2c2-157d4ebe-ad979cd4f296
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAccountConfigInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountConfigInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountConfigInfoResponseBody) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *GetAccountConfigInfoResponseBody) GetAvailableTime() *string {
	return s.AvailableTime
}

func (s *GetAccountConfigInfoResponseBody) GetBillQps() *string {
	return s.BillQps
}

func (s *GetAccountConfigInfoResponseBody) GetLadderType() *string {
	return s.LadderType
}

func (s *GetAccountConfigInfoResponseBody) GetMainAccountId() *string {
	return s.MainAccountId
}

func (s *GetAccountConfigInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountConfigInfoResponseBody) SetAccountStatus(v string) *GetAccountConfigInfoResponseBody {
	s.AccountStatus = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) SetAvailableTime(v string) *GetAccountConfigInfoResponseBody {
	s.AvailableTime = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) SetBillQps(v string) *GetAccountConfigInfoResponseBody {
	s.BillQps = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) SetLadderType(v string) *GetAccountConfigInfoResponseBody {
	s.LadderType = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) SetMainAccountId(v string) *GetAccountConfigInfoResponseBody {
	s.MainAccountId = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) SetRequestId(v string) *GetAccountConfigInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountConfigInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
