// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSensitiveDataResponseBody
	GetRequestId() *string
	SetSensitiveData(v map[string]interface{}) *GetSensitiveDataResponseBody
	GetSensitiveData() map[string]interface{}
}

type GetSensitiveDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the sensitive data returned. The information includes totalCount and sensDatas. sensDatas includes the following parameters:
	//
	// 	- guid: the ID of the metadata of the tenant. For example, the ID of the metadata in the MaxCompute compute engine is in the Project name.Table name.Column name format.
	//
	// 	- sensType: the type of the sensitive data.
	//
	// 	- sensLevel: the sensitivity level of the sensitive data
	SensitiveData map[string]interface{} `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty"`
}

func (s GetSensitiveDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSensitiveDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSensitiveDataResponseBody) GetSensitiveData() map[string]interface{} {
	return s.SensitiveData
}

func (s *GetSensitiveDataResponseBody) SetRequestId(v string) *GetSensitiveDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSensitiveDataResponseBody) SetSensitiveData(v map[string]interface{}) *GetSensitiveDataResponseBody {
	s.SensitiveData = v
	return s
}

func (s *GetSensitiveDataResponseBody) Validate() error {
	return dara.Validate(s)
}
