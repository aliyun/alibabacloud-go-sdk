// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceConfigurationSampleResponseBody
	GetRequestId() *string
	SetSample(v *GetResourceConfigurationSampleResponseBodySample) *GetResourceConfigurationSampleResponseBody
	GetSample() *GetResourceConfigurationSampleResponseBodySample
}

type GetResourceConfigurationSampleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB668487
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sample    *GetResourceConfigurationSampleResponseBodySample `json:"Sample,omitempty" xml:"Sample,omitempty" type:"Struct"`
}

func (s GetResourceConfigurationSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceConfigurationSampleResponseBody) GetSample() *GetResourceConfigurationSampleResponseBodySample {
	return s.Sample
}

func (s *GetResourceConfigurationSampleResponseBody) SetRequestId(v string) *GetResourceConfigurationSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBody) SetSample(v *GetResourceConfigurationSampleResponseBodySample) *GetResourceConfigurationSampleResponseBody {
	s.Sample = v
	return s
}

func (s *GetResourceConfigurationSampleResponseBody) Validate() error {
	if s.Sample != nil {
		if err := s.Sample.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceConfigurationSampleResponseBodySample struct {
	// example:
	//
	// 10093xxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// cn-hangzhou
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	Configuration    *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1646362560000
	ResourceCreationTime *string `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// example:
	//
	// 1
	ResourceDeleted *string `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// example:
	//
	// i-bp1aaegapahkh880x4wq
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// my-ecs
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// Running
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// {\\"key\\":\\"value\\"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 5
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetResourceConfigurationSampleResponseBodySample) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationSampleResponseBodySample) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetAccountId() *string {
	return s.AccountId
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetRegion() *string {
	return s.Region
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceCreationTime() *string {
	return s.ResourceCreationTime
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceDeleted() *string {
	return s.ResourceDeleted
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetTags() *string {
	return s.Tags
}

func (s *GetResourceConfigurationSampleResponseBodySample) GetVersion() *string {
	return s.Version
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetAccountId(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetAvailabilityZone(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetConfiguration(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.Configuration = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetRegion(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceCreationTime(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceDeleted(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceDeleted = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceId(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceName(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceStatus(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceStatus = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetResourceType(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetTags(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.Tags = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) SetVersion(v string) *GetResourceConfigurationSampleResponseBodySample {
	s.Version = &v
	return s
}

func (s *GetResourceConfigurationSampleResponseBodySample) Validate() error {
	return dara.Validate(s)
}
