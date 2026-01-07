// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentSubAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataAgentSubAccountInfoResponseBodyData) *GetDataAgentSubAccountInfoResponseBody
	GetData() *GetDataAgentSubAccountInfoResponseBodyData
	SetErrorCode(v string) *GetDataAgentSubAccountInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentSubAccountInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentSubAccountInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAgentSubAccountInfoResponseBody
	GetSuccess() *bool
}

type GetDataAgentSubAccountInfoResponseBody struct {
	Data *GetDataAgentSubAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 67E910F2-***-695C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAgentSubAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentSubAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentSubAccountInfoResponseBody) GetData() *GetDataAgentSubAccountInfoResponseBodyData {
	return s.Data
}

func (s *GetDataAgentSubAccountInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentSubAccountInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentSubAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentSubAccountInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAgentSubAccountInfoResponseBody) SetData(v *GetDataAgentSubAccountInfoResponseBodyData) *GetDataAgentSubAccountInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBody) SetErrorCode(v string) *GetDataAgentSubAccountInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBody) SetErrorMessage(v string) *GetDataAgentSubAccountInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBody) SetRequestId(v string) *GetDataAgentSubAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBody) SetSuccess(v bool) *GetDataAgentSubAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAgentSubAccountInfoResponseBodyData struct {
	// example:
	//
	// 1765960516
	CreateDate *int64 `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// example:
	//
	// 167*****166
	MainAccountId *string `json:"MainAccountId,omitempty" xml:"MainAccountId,omitempty"`
	// example:
	//
	// 1765962516
	UpdateDate *int64 `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// example:
	//
	// 20282*****7591
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// yunqitest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDataAgentSubAccountInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentSubAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) GetCreateDate() *int64 {
	return s.CreateDate
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) GetMainAccountId() *string {
	return s.MainAccountId
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) GetUpdateDate() *int64 {
	return s.UpdateDate
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) SetCreateDate(v int64) *GetDataAgentSubAccountInfoResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) SetMainAccountId(v string) *GetDataAgentSubAccountInfoResponseBodyData {
	s.MainAccountId = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) SetUpdateDate(v int64) *GetDataAgentSubAccountInfoResponseBodyData {
	s.UpdateDate = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) SetUserId(v string) *GetDataAgentSubAccountInfoResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) SetUserName(v string) *GetDataAgentSubAccountInfoResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetDataAgentSubAccountInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
