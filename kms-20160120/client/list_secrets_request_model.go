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
	// Specifies whether to return resource tags for each secret. Valid values:
	//
	// - `true`: Resource tags are returned.
	//
	// - `false` (default): Resource tags are not returned.
	//
	// example:
	//
	// false
	FetchTags *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	// Filters secrets based on specified conditions. The value is a list of up to 10 key-value pairs. When you filter by tag, the query returns a maximum of 4,000 resources. If more than 4,000 resources match the filter, call the [ListResourceTags](https://help.aliyun.com/document_detail/120090.html) operation.
	//
	// - Key
	//
	//   - Description: The filter property.
	//
	//   - Type: String.
	//
	// - Values
	//
	//   - Description: The filter value.
	//
	//   - Type: String.
	//
	//   - You can specify up to 10 items.
	//
	// Valid values for Key:
	//
	// - Set `Key` to **SecretName*	- to filter by secret name.
	//
	// - Set `Key` to **Description*	- to filter by secret description.
	//
	// - Set `Key` to **TagKey*	- to filter by tag key.
	//
	// - Set `Key` to **TagValue*	- to filter by tag value.
	//
	// - Set `Key` to **DKMSInstanceId*	- to filter by KMS instance ID.
	//
	// - Set Key to **SecretType*	- to filter by secret type. `Values` can be `Generic`, `RDS`, `Redis`, `RAMCredentials`, `ECS`, or `PolarDB`.
	//
	// - Set `Key` to **Creator*	- to filter by the creator of the secret.
	//
	// If you specify multiple values for a key, the filter applies a logical OR. For example, if you enter `[ {"Key":"SecretName", "Values":["sec1","sec2"]} ]`, this means: `(SecretName=sec1 OR SecretName=sec2)`.
	//
	// example:
	//
	// [{"Key":"SecretName", "Values":["Val1","Val2"]}]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The page number.<br>
	//
	// The value must be greater than 0.<br>
	//
	// Default: 1.<br><br>
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.<br>
	//
	// The value must be between 1 and 100.<br>
	//
	// Default: 10.<br><br>
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
