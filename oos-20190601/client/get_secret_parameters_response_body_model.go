// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvalidParameters(v []*string) *GetSecretParametersResponseBody
	GetInvalidParameters() []*string
	SetParameters(v []*GetSecretParametersResponseBodyParameters) *GetSecretParametersResponseBody
	GetParameters() []*GetSecretParametersResponseBodyParameters
	SetRequestId(v string) *GetSecretParametersResponseBody
	GetRequestId() *string
}

type GetSecretParametersResponseBody struct {
	// Invalid encryption parameter.
	InvalidParameters []*string `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	// The information about the encryption parameter.
	Parameters []*GetSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A5320F1D-92D9-44BB-A416-5FC525ED6D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBody) GetInvalidParameters() []*string {
	return s.InvalidParameters
}

func (s *GetSecretParametersResponseBody) GetParameters() []*GetSecretParametersResponseBodyParameters {
	return s.Parameters
}

func (s *GetSecretParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretParametersResponseBody) SetInvalidParameters(v []*string) *GetSecretParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetSecretParametersResponseBody) SetParameters(v []*GetSecretParametersResponseBodyParameters) *GetSecretParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *GetSecretParametersResponseBody) SetRequestId(v string) *GetSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParametersResponseBody) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSecretParametersResponseBodyParameters struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// {\\"AllowedValues\\": [\\"test\\"]}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-10-22T03:11:13Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// ssh-bp67acfmxazb4p****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-10-22T03:11:13Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// secretParameter,secretParameter1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBodyParameters) GetConstraints() *string {
	return s.Constraints
}

func (s *GetSecretParametersResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetSecretParametersResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetSecretParametersResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *GetSecretParametersResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *GetSecretParametersResponseBodyParameters) GetKeyId() *string {
	return s.KeyId
}

func (s *GetSecretParametersResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *GetSecretParametersResponseBodyParameters) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetSecretParametersResponseBodyParameters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSecretParametersResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *GetSecretParametersResponseBodyParameters) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetSecretParametersResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *GetSecretParametersResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetSecretParametersResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetSecretParametersResponseBodyParameters) GetValue() *string {
	return s.Value
}

func (s *GetSecretParametersResponseBodyParameters) SetConstraints(v string) *GetSecretParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetDescription(v string) *GetSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetId(v string) *GetSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetKeyId(v string) *GetSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetName(v string) *GetSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *GetSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetShareType(v string) *GetSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetType(v string) *GetSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetValue(v string) *GetSecretParametersResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
