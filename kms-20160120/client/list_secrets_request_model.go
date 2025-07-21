// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFetchTags(v string) *ListSecretsRequest
	GetFetchTags() *string
	SetFilters(v string) *ListSecretsRequest
	GetFilters() *string
	SetPageNumber(v int32) *ListSecretsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretsRequest
	GetPageSize() *int32
}

type ListSecretsRequest struct {
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// false
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// [{"Key":"SecretName", "Values":["Val1","Val2"]}]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The secret filter. The filter consists of one or more key-value pairs. You can specify one key-value pair or leave this parameter empty. If you use one tag key or tag value to filter resources, up to 4,000 resources can be queried. If you want to query more than 4,000 resources, call the [ListResourceTags](https://help.aliyun.com/document_detail/120090.html) operation.
	//
	// 	- Key
	//
	//     	- Description: the property that you want to filter.
	//
	//     	- Type: string.
	//
	//     	- Valid values:
	//
	//         	- SecretName: the secret name.
	//
	//         	- Description: the description of the secret.
	//
	//         	- TagKey: the tag key.
	//
	//         	- TagValue: the tag value.
	//
	// 	- Values
	//
	//     	- Description: the value to be included after filtering.
	//
	//     	- Type: string.
	//
	//     	- Length: 0 to 10.
	//
	//     	- Valid values:
	//
	//         	- If the Key field is set to SecretName, the value must be 1 to 192 characters in length and can contain letters, digits, and special characters `_ / + = . @ -`.
	//
	//         	- If the Key field is set to Description, the value must be 1 to 256 characters in length.
	//
	//         	- If the Key field is set to TagKey, the value must be 1 to 256 characters in length and can contain letters, digits, and special characters `/ _ - . + = @ :`.
	//
	//         	- If the Key field is set to TagValue, the value must be 1 to 256 characters in length and can contain letters, numbers, and special characters `/ _ - . + = @ :`.
	//
	// The logical relationship between values of the Values field in a key-value pair is OR. Example: `[ {"Key":"SecretName", "Values":["sec1","sec2"]}]`. In this example, the semantics are `SecretName=sec 1 OR SecretName=sec 2`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) GetFetchTags() *string {
	return s.FetchTags
}

func (s *ListSecretsRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListSecretsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretsRequest) SetFetchTags(v string) *ListSecretsRequest {
	s.FetchTags = &v
	return s
}

func (s *ListSecretsRequest) SetFilters(v string) *ListSecretsRequest {
	s.Filters = &v
	return s
}

func (s *ListSecretsRequest) SetPageNumber(v int32) *ListSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsRequest) SetPageSize(v int32) *ListSecretsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretsRequest) Validate() error {
	return dara.Validate(s)
}
