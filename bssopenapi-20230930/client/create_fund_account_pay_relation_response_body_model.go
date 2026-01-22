// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountPayRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*CreateFundAccountPayRelationResponseBodyData) *CreateFundAccountPayRelationResponseBody
	GetData() []*CreateFundAccountPayRelationResponseBodyData
	SetMetadata(v interface{}) *CreateFundAccountPayRelationResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CreateFundAccountPayRelationResponseBody
	GetRequestId() *string
}

type CreateFundAccountPayRelationResponseBody struct {
	Data []*CreateFundAccountPayRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFundAccountPayRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponseBody) GetData() []*CreateFundAccountPayRelationResponseBodyData {
	return s.Data
}

func (s *CreateFundAccountPayRelationResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateFundAccountPayRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFundAccountPayRelationResponseBody) SetData(v []*CreateFundAccountPayRelationResponseBodyData) *CreateFundAccountPayRelationResponseBody {
	s.Data = v
	return s
}

func (s *CreateFundAccountPayRelationResponseBody) SetMetadata(v interface{}) *CreateFundAccountPayRelationResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateFundAccountPayRelationResponseBody) SetRequestId(v string) *CreateFundAccountPayRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFundAccountPayRelationResponseBodyData struct {
	// example:
	//
	// 1501603440974415
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// Success
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// Successful
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateFundAccountPayRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateFundAccountPayRelationResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateFundAccountPayRelationResponseBodyData) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *CreateFundAccountPayRelationResponseBodyData) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateFundAccountPayRelationResponseBodyData) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetAccountId(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetAccountName(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetFundAccountId(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetResultCode(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.ResultCode = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetResultMessage(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.ResultMessage = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
