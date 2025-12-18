// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v string) *GetResourceTypePropertiesResponseBody
	GetConfiguration() *string
	SetRequestId(v string) *GetResourceTypePropertiesResponseBody
	GetRequestId() *string
}

type GetResourceTypePropertiesResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "description": "The ID of the resource group to which the instance belongs.\\n",
	//
	//       "title": "ResourceGroupId",
	//
	//       "type": "String",
	//
	//       "key": "ResourceGroupId",
	//
	//       "example": "rg-bp67acfmxazb4p****"
	//
	//     },
	//
	//     {
	//
	//       "description": "The billing method of the instance. Valid values:\\n\\n	- PrePaid: subscription\\n	- PostPaid: pay-as-you-go\\n",
	//
	//       "title": "InstanceChargeType",
	//
	//       "type": "String",
	//
	//       "key": "InstanceChargeType",
	//
	//       "example": "PostPaid"
	//
	//     }]
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C2868BF-47EE-5441-B34B-17F080B10DC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceTypePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypePropertiesResponseBody) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetResourceTypePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceTypePropertiesResponseBody) SetConfiguration(v string) *GetResourceTypePropertiesResponseBody {
	s.Configuration = &v
	return s
}

func (s *GetResourceTypePropertiesResponseBody) SetRequestId(v string) *GetResourceTypePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypePropertiesResponseBody) Validate() error {
	return dara.Validate(s)
}
