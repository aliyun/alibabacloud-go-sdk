// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConnectorModelsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListConnectorModelsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListConnectorModelsResponseBodyItems) *ListConnectorModelsResponseBody
	GetItems() []*ListConnectorModelsResponseBodyItems
	SetMaxResults(v int32) *ListConnectorModelsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListConnectorModelsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListConnectorModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConnectorModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConnectorModelsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListConnectorModelsResponseBody
	GetTotalCount() *int64
}

type ListConnectorModelsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of available models.
	Items []*ListConnectorModelsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of entries returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// dGVzdA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of models.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConnectorModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConnectorModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConnectorModelsResponseBody) GetItems() []*ListConnectorModelsResponseBodyItems {
	return s.Items
}

func (s *ListConnectorModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConnectorModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectorModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConnectorModelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConnectorModelsResponseBody) SetCode(v string) *ListConnectorModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetHttpStatusCode(v int32) *ListConnectorModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetItems(v []*ListConnectorModelsResponseBodyItems) *ListConnectorModelsResponseBody {
	s.Items = v
	return s
}

func (s *ListConnectorModelsResponseBody) SetMaxResults(v int32) *ListConnectorModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetMessage(v string) *ListConnectorModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetNextToken(v string) *ListConnectorModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetRequestId(v string) *ListConnectorModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetSuccess(v bool) *ListConnectorModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConnectorModelsResponseBody) SetTotalCount(v int64) *ListConnectorModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectorModelsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConnectorModelsResponseBodyItems struct {
	// The description of the model.
	//
	// example:
	//
	// 通义千问旗舰模型
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the model.
	//
	// example:
	//
	// Qwen3 Max
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Key ID
	//
	// example:
	//
	// ckey-xxxx
	KeyId *string `json:"keyId,omitempty" xml:"keyId,omitempty"`
	// The key name.
	//
	// example:
	//
	// default
	KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
	// The stable identifier of the model.
	//
	// example:
	//
	// qwen3-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The source of the model. Valid values:
	//
	// - official: an official model.
	//
	// - enterprise: an enterprise-specific model.
	//
	// example:
	//
	// official
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListConnectorModelsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorModelsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListConnectorModelsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListConnectorModelsResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListConnectorModelsResponseBodyItems) GetKeyId() *string {
	return s.KeyId
}

func (s *ListConnectorModelsResponseBodyItems) GetKeyName() *string {
	return s.KeyName
}

func (s *ListConnectorModelsResponseBodyItems) GetModelId() *string {
	return s.ModelId
}

func (s *ListConnectorModelsResponseBodyItems) GetSource() *string {
	return s.Source
}

func (s *ListConnectorModelsResponseBodyItems) SetDescription(v string) *ListConnectorModelsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) SetDisplayName(v string) *ListConnectorModelsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) SetKeyId(v string) *ListConnectorModelsResponseBodyItems {
	s.KeyId = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) SetKeyName(v string) *ListConnectorModelsResponseBodyItems {
	s.KeyName = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) SetModelId(v string) *ListConnectorModelsResponseBodyItems {
	s.ModelId = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) SetSource(v string) *ListConnectorModelsResponseBodyItems {
	s.Source = &v
	return s
}

func (s *ListConnectorModelsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
