// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListApiKeysResponseBodyItems) *ListApiKeysResponseBody
	GetItems() []*ListApiKeysResponseBodyItems
	SetMaxResults(v int32) *ListApiKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiKeysResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListApiKeysResponseBody
	GetTotalRecordCount() *int32
}

type ListApiKeysResponseBody struct {
	// The list of API keys.
	Items []*ListApiKeysResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page in a paged query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBody) GetItems() []*ListApiKeysResponseBodyItems {
	return s.Items
}

func (s *ListApiKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiKeysResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListApiKeysResponseBody) SetItems(v []*ListApiKeysResponseBodyItems) *ListApiKeysResponseBody {
	s.Items = v
	return s
}

func (s *ListApiKeysResponseBody) SetMaxResults(v int32) *ListApiKeysResponseBody {
	s.MaxResults = &v
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

func (s *ListApiKeysResponseBody) SetTotalRecordCount(v int32) *ListApiKeysResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListApiKeysResponseBody) Validate() error {
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

type ListApiKeysResponseBodyItems struct {
	// The service IDs.
	AuthServices []*ListApiKeysResponseBodyItemsAuthServices `json:"AuthServices,omitempty" xml:"AuthServices,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the API key.
	//
	// example:
	//
	// my first api key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API key.
	//
	// example:
	//
	// api-xxxxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// my first api key
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// sk-12345****
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
}

func (s ListApiKeysResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyItems) GetAuthServices() []*ListApiKeysResponseBodyItemsAuthServices {
	return s.AuthServices
}

func (s *ListApiKeysResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApiKeysResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListApiKeysResponseBodyItems) GetKeyId() *string {
	return s.KeyId
}

func (s *ListApiKeysResponseBodyItems) GetKeyName() *string {
	return s.KeyName
}

func (s *ListApiKeysResponseBodyItems) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *ListApiKeysResponseBodyItems) SetAuthServices(v []*ListApiKeysResponseBodyItemsAuthServices) *ListApiKeysResponseBodyItems {
	s.AuthServices = v
	return s
}

func (s *ListApiKeysResponseBodyItems) SetCreateTime(v string) *ListApiKeysResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListApiKeysResponseBodyItems) SetDescription(v string) *ListApiKeysResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListApiKeysResponseBodyItems) SetKeyId(v string) *ListApiKeysResponseBodyItems {
	s.KeyId = &v
	return s
}

func (s *ListApiKeysResponseBodyItems) SetKeyName(v string) *ListApiKeysResponseBodyItems {
	s.KeyName = &v
	return s
}

func (s *ListApiKeysResponseBodyItems) SetKeyPrefix(v string) *ListApiKeysResponseBodyItems {
	s.KeyPrefix = &v
	return s
}

func (s *ListApiKeysResponseBodyItems) Validate() error {
	if s.AuthServices != nil {
		for _, item := range s.AuthServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiKeysResponseBodyItemsAuthServices struct {
	// The service IDs.
	//
	// example:
	//
	// agdb-xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type.
	//
	// Valid values:
	//
	// - memory
	//
	// - drama
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListApiKeysResponseBodyItemsAuthServices) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyItemsAuthServices) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyItemsAuthServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListApiKeysResponseBodyItemsAuthServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListApiKeysResponseBodyItemsAuthServices) SetServiceId(v string) *ListApiKeysResponseBodyItemsAuthServices {
	s.ServiceId = &v
	return s
}

func (s *ListApiKeysResponseBodyItemsAuthServices) SetServiceType(v string) *ListApiKeysResponseBodyItemsAuthServices {
	s.ServiceType = &v
	return s
}

func (s *ListApiKeysResponseBodyItemsAuthServices) Validate() error {
	return dara.Validate(s)
}
