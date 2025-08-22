// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) *BatchSetDcdnDomainConfigsResponseBody
	GetDomainConfigList() *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList
	SetRequestId(v string) *BatchSetDcdnDomainConfigsResponseBody
	GetRequestId() *string
}

type BatchSetDcdnDomainConfigsResponseBody struct {
	// The list of domain configurations.
	DomainConfigList *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDcdnDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponseBody) GetDomainConfigList() *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *BatchSetDcdnDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetDcdnDomainConfigsResponseBody) SetDomainConfigList(v *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) *BatchSetDcdnDomainConfigsResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBody) SetRequestId(v string) *BatchSetDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchSetDcdnDomainConfigsResponseBodyDomainConfigList struct {
	DomainConfigModel []*BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel `json:"DomainConfigModel,omitempty" xml:"DomainConfigModel,omitempty" type:"Repeated"`
}

func (s BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) GetDomainConfigModel() []*BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	return s.DomainConfigModel
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) SetDomainConfigModel(v []*BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList {
	s.DomainConfigModel = v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigList) Validate() error {
	return dara.Validate(s)
}

type BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 123456
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// set_resp_header
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) GetFunctionName() *string {
	return s.FunctionName
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetConfigId(v int64) *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.ConfigId = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetDomainName(v string) *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.DomainName = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetFunctionName(v string) *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.FunctionName = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) Validate() error {
	return dara.Validate(s)
}
