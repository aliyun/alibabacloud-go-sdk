// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTokens(v []*ListApplicationTokensResponseBodyApplicationTokens) *ListApplicationTokensResponseBody
	GetApplicationTokens() []*ListApplicationTokensResponseBodyApplicationTokens
	SetRequestId(v string) *ListApplicationTokensResponseBody
	GetRequestId() *string
}

type ListApplicationTokensResponseBody struct {
	ApplicationTokens []*ListApplicationTokensResponseBodyApplicationTokens `json:"ApplicationTokens,omitempty" xml:"ApplicationTokens,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationTokensResponseBody) GetApplicationTokens() []*ListApplicationTokensResponseBodyApplicationTokens {
	return s.ApplicationTokens
}

func (s *ListApplicationTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationTokensResponseBody) SetApplicationTokens(v []*ListApplicationTokensResponseBodyApplicationTokens) *ListApplicationTokensResponseBody {
	s.ApplicationTokens = v
	return s
}

func (s *ListApplicationTokensResponseBody) SetRequestId(v string) *ListApplicationTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationTokensResponseBody) Validate() error {
	if s.ApplicationTokens != nil {
		for _, item := range s.ApplicationTokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationTokensResponseBodyApplicationTokens struct {
	// aliUid。
	//
	// example:
	//
	// 1973166921975xxxx
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用ID
	//
	// example:
	//
	// app_m7ar5tms4dwtggavalk3j3mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用token
	//
	// example:
	//
	// SATFwqX8zxGf83pJcJw78KFGjmrft4erWeZYBGS8oE7NN6qoE217yaJpUdMb1UuuGqhDiF43sCA4CF91CTL5iGntqwyLuaAcS9FJ9HfGadE5a7TjiwVafwrBYktxxxx
	ApplicationToken *string `json:"ApplicationToken,omitempty" xml:"ApplicationToken,omitempty"`
	// 应用token ID
	//
	// example:
	//
	// token_ndfxxigahelfne2y2hodehrxxxx
	ApplicationTokenId *string `json:"ApplicationTokenId,omitempty" xml:"ApplicationTokenId,omitempty"`
	// 应用token类型
	//
	// example:
	//
	// bearer_token
	ApplicationTokenType *string `json:"ApplicationTokenType,omitempty" xml:"ApplicationTokenType,omitempty"`
	// example:
	//
	// 1747796654000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用token描述
	//
	// example:
	//
	// jwqtts-0430
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 到期时间
	//
	// example:
	//
	// 1747796654000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ki6hd7ihir4ybawogqk6kqxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 最后使用时间
	//
	// example:
	//
	// 1747796654000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// 应用状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationTokensResponseBodyApplicationTokens) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTokensResponseBodyApplicationTokens) GoString() string {
	return s.String()
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetApplicationToken() *string {
	return s.ApplicationToken
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetApplicationTokenType() *string {
	return s.ApplicationTokenType
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetAliUid(v int64) *ListApplicationTokensResponseBodyApplicationTokens {
	s.AliUid = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetApplicationId(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetApplicationToken(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.ApplicationToken = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetApplicationTokenId(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.ApplicationTokenId = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetApplicationTokenType(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.ApplicationTokenType = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetCreateTime(v int64) *ListApplicationTokensResponseBodyApplicationTokens {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetDescription(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.Description = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetExpirationTime(v int64) *ListApplicationTokensResponseBodyApplicationTokens {
	s.ExpirationTime = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetInstanceId(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetLastUsedTime(v int64) *ListApplicationTokensResponseBodyApplicationTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) SetStatus(v string) *ListApplicationTokensResponseBodyApplicationTokens {
	s.Status = &v
	return s
}

func (s *ListApplicationTokensResponseBodyApplicationTokens) Validate() error {
	return dara.Validate(s)
}
