// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ListModelOperatorApiKeysResponseBodyApiKeys) *ListModelOperatorApiKeysResponseBody
	GetApiKeys() []*ListModelOperatorApiKeysResponseBodyApiKeys
	SetPageNumber(v int32) *ListModelOperatorApiKeysResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListModelOperatorApiKeysResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListModelOperatorApiKeysResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListModelOperatorApiKeysResponseBody
	GetTotalRecordCount() *int32
}

type ListModelOperatorApiKeysResponseBody struct {
	ApiKeys []*ListModelOperatorApiKeysResponseBodyApiKeys `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListModelOperatorApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelOperatorApiKeysResponseBody) GetApiKeys() []*ListModelOperatorApiKeysResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *ListModelOperatorApiKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelOperatorApiKeysResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListModelOperatorApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelOperatorApiKeysResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListModelOperatorApiKeysResponseBody) SetApiKeys(v []*ListModelOperatorApiKeysResponseBodyApiKeys) *ListModelOperatorApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *ListModelOperatorApiKeysResponseBody) SetPageNumber(v int32) *ListModelOperatorApiKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBody) SetPageRecordCount(v int32) *ListModelOperatorApiKeysResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBody) SetRequestId(v string) *ListModelOperatorApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBody) SetTotalRecordCount(v int32) *ListModelOperatorApiKeysResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBody) Validate() error {
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

type ListModelOperatorApiKeysResponseBodyApiKeys struct {
	// example:
	//
	// 1
	ApiKeyId *int32 `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// 2026-06-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test-apikey
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://xxxx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s ListModelOperatorApiKeysResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorApiKeysResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) GetApiKeyId() *int32 {
	return s.ApiKeyId
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) GetDescription() *string {
	return s.Description
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) SetApiKeyId(v int32) *ListModelOperatorApiKeysResponseBodyApiKeys {
	s.ApiKeyId = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) SetCreateTime(v string) *ListModelOperatorApiKeysResponseBodyApiKeys {
	s.CreateTime = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) SetDescription(v string) *ListModelOperatorApiKeysResponseBodyApiKeys {
	s.Description = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) SetEndpoint(v string) *ListModelOperatorApiKeysResponseBodyApiKeys {
	s.Endpoint = &v
	return s
}

func (s *ListModelOperatorApiKeysResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}
