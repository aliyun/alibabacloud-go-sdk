// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*DescribeAIDBClusterApiKeysResponseBodyApiKeys) *DescribeAIDBClusterApiKeysResponseBody
	GetApiKeys() []*DescribeAIDBClusterApiKeysResponseBodyApiKeys
	SetPageNumber(v string) *DescribeAIDBClusterApiKeysResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAIDBClusterApiKeysResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeAIDBClusterApiKeysResponseBody
	GetRequestId() *string
}

type DescribeAIDBClusterApiKeysResponseBody struct {
	// API Keys。
	ApiKeys    []*DescribeAIDBClusterApiKeysResponseBodyApiKeys `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	PageNumber *string                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F2007D3-7E74-4ECB-89A8-BF130D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIDBClusterApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterApiKeysResponseBody) GetApiKeys() []*DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *DescribeAIDBClusterApiKeysResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAIDBClusterApiKeysResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAIDBClusterApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterApiKeysResponseBody) SetApiKeys(v []*DescribeAIDBClusterApiKeysResponseBodyApiKeys) *DescribeAIDBClusterApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBody) SetPageNumber(v string) *DescribeAIDBClusterApiKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBody) SetPageSize(v string) *DescribeAIDBClusterApiKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBody) SetRequestId(v string) *DescribeAIDBClusterApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBody) Validate() error {
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

type DescribeAIDBClusterApiKeysResponseBodyApiKeys struct {
	// The API key of the model service.
	//
	// example:
	//
	// Scxxx-xxx-x-xxWW
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-09T03:19:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the API key.
	//
	// example:
	//
	// my api key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ApiKey ID
	//
	// example:
	//
	// 573
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the API key.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAIDBClusterApiKeysResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterApiKeysResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) GetDescription() *string {
	return s.Description
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) GetId() *string {
	return s.Id
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) GetStatus() *string {
	return s.Status
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) SetApiKey(v string) *DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	s.ApiKey = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) SetCreateTime(v string) *DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) SetDescription(v string) *DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	s.Description = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) SetId(v string) *DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	s.Id = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) SetStatus(v string) *DescribeAIDBClusterApiKeysResponseBodyApiKeys {
	s.Status = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}
