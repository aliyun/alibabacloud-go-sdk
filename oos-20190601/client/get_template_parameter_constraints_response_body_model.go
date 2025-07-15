// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParameterConstraintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterConstraints(v map[string]interface{}) *GetTemplateParameterConstraintsResponseBody
	GetParameterConstraints() map[string]interface{}
	SetRequestId(v string) *GetTemplateParameterConstraintsResponseBody
	GetRequestId() *string
}

type GetTemplateParameterConstraintsResponseBody struct {
	// The constraints of the parameters.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "Type": "String",
	//
	//       "AllowedValues": [
	//
	//         "ecs.n1.tiny",
	//
	//         "ecs.r8a.4xlarge",
	//
	//         "ecs.n2.xlarge",
	//
	//         "ecs.c7.2xlarge",
	//
	//         "ecs.c8i.4xlarge",
	//
	//         "ecs.g8i.48xlarge",
	//
	//         "ecs.c8a.4xlarge",
	//
	//         "ecs.i2.4xlarge",
	//
	//         "ecs.r8y.2xlarge"
	//
	//       ],
	//
	//       "AssociationParameterNames": [
	//
	//         "RegionId",
	//
	//         "zoneId"
	//
	//       ],
	//
	//       "ParameterKey": "instanceType"
	//
	//     },
	//
	//     {
	//
	//       "Type": "String",
	//
	//       "AllowedValues": [],
	//
	//       "AssociationParameterNames": [
	//
	//         "RegionId",
	//
	//         "zoneId",
	//
	//         "InstanceType"
	//
	//       ],
	//
	//       "ParameterKey": "systemDiskCategory"
	//
	//     }
	//
	//   ]
	ParameterConstraints map[string]interface{} `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBEC8072-BEC2-478E-8EAE-E723BA79CF19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBody) GetParameterConstraints() map[string]interface{} {
	return s.ParameterConstraints
}

func (s *GetTemplateParameterConstraintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateParameterConstraintsResponseBody) SetParameterConstraints(v map[string]interface{}) *GetTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBody) Validate() error {
	return dara.Validate(s)
}
