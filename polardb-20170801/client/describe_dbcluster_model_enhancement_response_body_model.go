// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterModelEnhancementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeDBClusterModelEnhancementResponseBody
	GetCreateTime() *string
	SetModelEnhancementKeyNumber(v int32) *DescribeDBClusterModelEnhancementResponseBody
	GetModelEnhancementKeyNumber() *int32
	SetModelEnhancementKeys(v []*DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) *DescribeDBClusterModelEnhancementResponseBody
	GetModelEnhancementKeys() []*DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys
	SetModelEnhancementStatus(v string) *DescribeDBClusterModelEnhancementResponseBody
	GetModelEnhancementStatus() *string
	SetRequestId(v string) *DescribeDBClusterModelEnhancementResponseBody
	GetRequestId() *string
}

type DescribeDBClusterModelEnhancementResponseBody struct {
	// The time when the model enhancement feature was enabled, in the ISO 8601 UTC format (YYYY-MM-DDThh:mm:ssZ).
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of API keys for model enhancement.
	//
	// example:
	//
	// 1
	ModelEnhancementKeyNumber *int32 `json:"ModelEnhancementKeyNumber,omitempty" xml:"ModelEnhancementKeyNumber,omitempty"`
	// The list of API keys for model enhancement.
	ModelEnhancementKeys []*DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys `json:"ModelEnhancementKeys,omitempty" xml:"ModelEnhancementKeys,omitempty" type:"Repeated"`
	// The status of the model enhancement feature. A value of ON indicates that the feature is enabled.
	//
	// example:
	//
	// ON
	ModelEnhancementStatus *string `json:"ModelEnhancementStatus,omitempty" xml:"ModelEnhancementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterModelEnhancementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterModelEnhancementResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterModelEnhancementResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClusterModelEnhancementResponseBody) GetModelEnhancementKeyNumber() *int32 {
	return s.ModelEnhancementKeyNumber
}

func (s *DescribeDBClusterModelEnhancementResponseBody) GetModelEnhancementKeys() []*DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	return s.ModelEnhancementKeys
}

func (s *DescribeDBClusterModelEnhancementResponseBody) GetModelEnhancementStatus() *string {
	return s.ModelEnhancementStatus
}

func (s *DescribeDBClusterModelEnhancementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterModelEnhancementResponseBody) SetCreateTime(v string) *DescribeDBClusterModelEnhancementResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBody) SetModelEnhancementKeyNumber(v int32) *DescribeDBClusterModelEnhancementResponseBody {
	s.ModelEnhancementKeyNumber = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBody) SetModelEnhancementKeys(v []*DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) *DescribeDBClusterModelEnhancementResponseBody {
	s.ModelEnhancementKeys = v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBody) SetModelEnhancementStatus(v string) *DescribeDBClusterModelEnhancementResponseBody {
	s.ModelEnhancementStatus = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBody) SetRequestId(v string) *DescribeDBClusterModelEnhancementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBody) Validate() error {
	if s.ModelEnhancementKeys != nil {
		for _, item := range s.ModelEnhancementKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys struct {
	// The value of the API key.
	//
	// example:
	//
	// ak-12345
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The ID of the API key.
	//
	// example:
	//
	// key-001
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The time when the API key was created, in the ISO 8601 UTC format (YYYY-MM-DDThh:mm:ssZ).
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the API key.
	//
	// example:
	//
	// Model enhancement key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time of the API key, in the ISO 8601 UTC format (YYYY-MM-DDThh:mm:ssZ).
	//
	// example:
	//
	// 2025-12-31T23:59:59Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) SetApiKey(v string) *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	s.ApiKey = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) SetApiKeyId(v string) *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) SetCreateTime(v string) *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) SetDescription(v string) *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	s.Description = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) SetExpireTime(v string) *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementResponseBodyModelEnhancementKeys) Validate() error {
	return dara.Validate(s)
}
