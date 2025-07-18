// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSupportFeaturesResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *DescribeSupportFeaturesResponseBody
	GetRequestId() *string
	SetSupportFeatureList(v string) *DescribeSupportFeaturesResponseBody
	GetSupportFeatureList() *string
}

type DescribeSupportFeaturesResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 63E5BE60-91FF-57F1-B873-7F1EB734B93D_2724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The features supported by the instance. Valid values:
	//
	// 	- sample_data: sample dataset. For more information, see [Sample dataset](https://help.aliyun.com/document_detail/452278.html).
	//
	// 	- diagnose_and_optimize: diagnostics and optimization. For more information, see [Diagnostics and optimization](https://help.aliyun.com/document_detail/323453.html).
	//
	// example:
	//
	// [ "sample_data", "diagnose_and_optimize" ]
	SupportFeatureList *string `json:"SupportFeatureList,omitempty" xml:"SupportFeatureList,omitempty"`
}

func (s DescribeSupportFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSupportFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportFeaturesResponseBody) GetSupportFeatureList() *string {
	return s.SupportFeatureList
}

func (s *DescribeSupportFeaturesResponseBody) SetDBInstanceId(v string) *DescribeSupportFeaturesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetRequestId(v string) *DescribeSupportFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) SetSupportFeatureList(v string) *DescribeSupportFeaturesResponseBody {
	s.SupportFeatureList = &v
	return s
}

func (s *DescribeSupportFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}
