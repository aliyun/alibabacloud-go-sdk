// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetCdnDomainConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigList(v *BatchSetCdnDomainConfigResponseBodyDomainConfigList) *BatchSetCdnDomainConfigResponseBody
	GetDomainConfigList() *BatchSetCdnDomainConfigResponseBodyDomainConfigList
	SetRequestId(v string) *BatchSetCdnDomainConfigResponseBody
	GetRequestId() *string
}

type BatchSetCdnDomainConfigResponseBody struct {
	// The list of domain configurations.
	DomainConfigList *BatchSetCdnDomainConfigResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetCdnDomainConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponseBody) GetDomainConfigList() *BatchSetCdnDomainConfigResponseBodyDomainConfigList {
	return s.DomainConfigList
}

func (s *BatchSetCdnDomainConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSetCdnDomainConfigResponseBody) SetDomainConfigList(v *BatchSetCdnDomainConfigResponseBodyDomainConfigList) *BatchSetCdnDomainConfigResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBody) SetRequestId(v string) *BatchSetCdnDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchSetCdnDomainConfigResponseBodyDomainConfigList struct {
	DomainConfigModel []*BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel `json:"DomainConfigModel,omitempty" xml:"DomainConfigModel,omitempty" type:"Repeated"`
}

func (s BatchSetCdnDomainConfigResponseBodyDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigList) GetDomainConfigModel() []*BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel {
	return s.DomainConfigModel
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigList) SetDomainConfigModel(v []*BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) *BatchSetCdnDomainConfigResponseBodyDomainConfigList {
	s.DomainConfigModel = v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigList) Validate() error {
	return dara.Validate(s)
}

type BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 1234567
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

func (s BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) String() string {
	return dara.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) GetFunctionName() *string {
	return s.FunctionName
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) SetConfigId(v int64) *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel {
	s.ConfigId = &v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) SetDomainName(v string) *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel {
	s.DomainName = &v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) SetFunctionName(v string) *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel {
	s.FunctionName = &v
	return s
}

func (s *BatchSetCdnDomainConfigResponseBodyDomainConfigListDomainConfigModel) Validate() error {
	return dara.Validate(s)
}
