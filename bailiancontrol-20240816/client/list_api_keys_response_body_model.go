// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ListApiKeysResponseBodyApiKeys) *ListApiKeysResponseBody
	GetApiKeys() []*ListApiKeysResponseBodyApiKeys
	SetCode(v string) *ListApiKeysResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListApiKeysResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListApiKeysResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListApiKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApiKeysResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListApiKeysResponseBody
	GetTotalCount() *int32
}

type ListApiKeysResponseBody struct {
	// apiKey
	ApiKeys        []*ListApiKeysResponseBodyApiKeys `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	MaxResults     *int32                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	NextToken      *string                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBody) GetApiKeys() []*ListApiKeysResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *ListApiKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApiKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiKeysResponseBody) SetApiKeys(v []*ListApiKeysResponseBodyApiKeys) *ListApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *ListApiKeysResponseBody) SetCode(v string) *ListApiKeysResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiKeysResponseBody) SetHttpStatusCode(v int32) *ListApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApiKeysResponseBody) SetMaxResults(v int32) *ListApiKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApiKeysResponseBody) SetMessage(v string) *ListApiKeysResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiKeysResponseBody) SetNextToken(v string) *ListApiKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApiKeysResponseBody) SetRequestId(v string) *ListApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiKeysResponseBody) SetSuccess(v bool) *ListApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiKeysResponseBody) SetTotalCount(v int32) *ListApiKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiKeysResponseBody) Validate() error {
	if s.ApiKeys != nil {
		for _, item := range s.ApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiKeysResponseBodyApiKeys struct {
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	ApikeyId    *string `json:"apikeyId,omitempty" xml:"apikeyId,omitempty"`
	Blocked     *int32  `json:"blocked,omitempty" xml:"blocked,omitempty"`
	CreateTime  *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ExpireTime  *int64  `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	ExtData     *string `json:"extData,omitempty" xml:"extData,omitempty"`
	ParentUid   *string `json:"parentUid,omitempty" xml:"parentUid,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListApiKeysResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyApiKeys) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ListApiKeysResponseBodyApiKeys) GetApikeyId() *string {
	return s.ApikeyId
}

func (s *ListApiKeysResponseBodyApiKeys) GetBlocked() *int32 {
	return s.Blocked
}

func (s *ListApiKeysResponseBodyApiKeys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApiKeysResponseBodyApiKeys) GetCreator() *string {
	return s.Creator
}

func (s *ListApiKeysResponseBodyApiKeys) GetDescription() *string {
	return s.Description
}

func (s *ListApiKeysResponseBodyApiKeys) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListApiKeysResponseBodyApiKeys) GetExtData() *string {
	return s.ExtData
}

func (s *ListApiKeysResponseBodyApiKeys) GetParentUid() *string {
	return s.ParentUid
}

func (s *ListApiKeysResponseBodyApiKeys) GetType() *int32 {
	return s.Type
}

func (s *ListApiKeysResponseBodyApiKeys) GetUid() *string {
	return s.Uid
}

func (s *ListApiKeysResponseBodyApiKeys) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysResponseBodyApiKeys) SetApiKeyValue(v string) *ListApiKeysResponseBodyApiKeys {
	s.ApiKeyValue = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetApikeyId(v string) *ListApiKeysResponseBodyApiKeys {
	s.ApikeyId = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetBlocked(v int32) *ListApiKeysResponseBodyApiKeys {
	s.Blocked = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetCreateTime(v int64) *ListApiKeysResponseBodyApiKeys {
	s.CreateTime = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetCreator(v string) *ListApiKeysResponseBodyApiKeys {
	s.Creator = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetDescription(v string) *ListApiKeysResponseBodyApiKeys {
	s.Description = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetExpireTime(v int64) *ListApiKeysResponseBodyApiKeys {
	s.ExpireTime = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetExtData(v string) *ListApiKeysResponseBodyApiKeys {
	s.ExtData = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetParentUid(v string) *ListApiKeysResponseBodyApiKeys {
	s.ParentUid = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetType(v int32) *ListApiKeysResponseBodyApiKeys {
	s.Type = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetUid(v string) *ListApiKeysResponseBodyApiKeys {
	s.Uid = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetWorkspaceId(v string) *ListApiKeysResponseBodyApiKeys {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}
