// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedeemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRedeemResponseBody
	GetCode() *string
	SetData(v *QueryRedeemResponseBodyData) *QueryRedeemResponseBody
	GetData() *QueryRedeemResponseBodyData
	SetMessage(v string) *QueryRedeemResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRedeemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRedeemResponseBody
	GetSuccess() *bool
}

type QueryRedeemResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryRedeemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E503DC7B-E4F0-4B3C-BC89-BCECF1338F0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRedeemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRedeemResponseBody) GetData() *QueryRedeemResponseBodyData {
	return s.Data
}

func (s *QueryRedeemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRedeemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRedeemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRedeemResponseBody) SetCode(v string) *QueryRedeemResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRedeemResponseBody) SetData(v *QueryRedeemResponseBodyData) *QueryRedeemResponseBody {
	s.Data = v
	return s
}

func (s *QueryRedeemResponseBody) SetMessage(v string) *QueryRedeemResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRedeemResponseBody) SetRequestId(v string) *QueryRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRedeemResponseBody) SetSuccess(v bool) *QueryRedeemResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRedeemResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRedeemResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Redeem   *QueryRedeemResponseBodyDataRedeem `json:"Redeem,omitempty" xml:"Redeem,omitempty" type:"Struct"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRedeemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *QueryRedeemResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryRedeemResponseBodyData) GetRedeem() *QueryRedeemResponseBodyDataRedeem {
	return s.Redeem
}

func (s *QueryRedeemResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryRedeemResponseBodyData) SetPageNum(v int64) *QueryRedeemResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRedeemResponseBodyData) SetPageSize(v int64) *QueryRedeemResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRedeemResponseBodyData) SetRedeem(v *QueryRedeemResponseBodyDataRedeem) *QueryRedeemResponseBodyData {
	s.Redeem = v
	return s
}

func (s *QueryRedeemResponseBodyData) SetTotalCount(v int64) *QueryRedeemResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRedeemResponseBodyData) Validate() error {
	if s.Redeem != nil {
		if err := s.Redeem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRedeemResponseBodyDataRedeem struct {
	Redeem []*QueryRedeemResponseBodyDataRedeemRedeem `json:"Redeem,omitempty" xml:"Redeem,omitempty" type:"Repeated"`
}

func (s QueryRedeemResponseBodyDataRedeem) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemResponseBodyDataRedeem) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyDataRedeem) GetRedeem() []*QueryRedeemResponseBodyDataRedeemRedeem {
	return s.Redeem
}

func (s *QueryRedeemResponseBodyDataRedeem) SetRedeem(v []*QueryRedeemResponseBodyDataRedeemRedeem) *QueryRedeemResponseBodyDataRedeem {
	s.Redeem = v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeem) Validate() error {
	if s.Redeem != nil {
		for _, item := range s.Redeem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRedeemResponseBodyDataRedeemRedeem struct {
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	Balance            *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	EffectiveTime      *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime         *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	GrantedTime        *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	NominalValue       *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	RedeemId           *string `json:"RedeemId,omitempty" xml:"RedeemId,omitempty"`
	RedeemNo           *string `json:"RedeemNo,omitempty" xml:"RedeemNo,omitempty"`
	Specification      *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryRedeemResponseBodyDataRedeemRedeem) String() string {
	return dara.Prettify(s)
}

func (s QueryRedeemResponseBodyDataRedeemRedeem) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetBalance() *string {
	return s.Balance
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetGrantedTime() *string {
	return s.GrantedTime
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetNominalValue() *string {
	return s.NominalValue
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetRedeemId() *string {
	return s.RedeemId
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetRedeemNo() *string {
	return s.RedeemNo
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetSpecification() *string {
	return s.Specification
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) GetStatus() *string {
	return s.Status
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetApplicableProducts(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetBalance(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Balance = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetEffectiveTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.EffectiveTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetExpiryTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.ExpiryTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetGrantedTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.GrantedTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetNominalValue(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.NominalValue = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetRedeemId(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.RedeemId = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetRedeemNo(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.RedeemNo = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetSpecification(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Specification = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetStatus(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Status = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) Validate() error {
	return dara.Validate(s)
}
