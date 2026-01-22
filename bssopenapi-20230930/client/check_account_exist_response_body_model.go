// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckAccountExistResponseBodyData) *CheckAccountExistResponseBody
	GetData() *CheckAccountExistResponseBodyData
	SetRequestId(v string) *CheckAccountExistResponseBody
	GetRequestId() *string
}

type CheckAccountExistResponseBody struct {
	Data *CheckAccountExistResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountExistResponseBody) GetData() *CheckAccountExistResponseBodyData {
	return s.Data
}

func (s *CheckAccountExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAccountExistResponseBody) SetData(v *CheckAccountExistResponseBodyData) *CheckAccountExistResponseBody {
	s.Data = v
	return s
}

func (s *CheckAccountExistResponseBody) SetRequestId(v string) *CheckAccountExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountExistResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckAccountExistResponseBodyData struct {
	// example:
	//
	// 1990699401005016
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 01
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 1904635303738977
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckAccountExistResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountExistResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckAccountExistResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *CheckAccountExistResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *CheckAccountExistResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *CheckAccountExistResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *CheckAccountExistResponseBodyData) SetAccountId(v string) *CheckAccountExistResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *CheckAccountExistResponseBodyData) SetNickName(v string) *CheckAccountExistResponseBodyData {
	s.NickName = &v
	return s
}

func (s *CheckAccountExistResponseBodyData) SetPk(v string) *CheckAccountExistResponseBodyData {
	s.Pk = &v
	return s
}

func (s *CheckAccountExistResponseBodyData) SetResult(v bool) *CheckAccountExistResponseBodyData {
	s.Result = &v
	return s
}

func (s *CheckAccountExistResponseBodyData) Validate() error {
	return dara.Validate(s)
}
