// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNacosGrayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryNacosGrayConfigResponseBodyData) *QueryNacosGrayConfigResponseBody
	GetData() *QueryNacosGrayConfigResponseBodyData
	SetRequestId(v string) *QueryNacosGrayConfigResponseBody
	GetRequestId() *string
}

type QueryNacosGrayConfigResponseBody struct {
	Data *QueryNacosGrayConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNacosGrayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNacosGrayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNacosGrayConfigResponseBody) GetData() *QueryNacosGrayConfigResponseBodyData {
	return s.Data
}

func (s *QueryNacosGrayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNacosGrayConfigResponseBody) SetData(v *QueryNacosGrayConfigResponseBodyData) *QueryNacosGrayConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryNacosGrayConfigResponseBody) SetRequestId(v string) *QueryNacosGrayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryNacosGrayConfigResponseBodyData struct {
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// text
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// nacos.config.gray.label=test
	GrayRule *string `json:"GrayRule,omitempty" xml:"GrayRule,omitempty"`
	// example:
	//
	// Beta(IP)
	GrayRuleName *string `json:"GrayRuleName,omitempty" xml:"GrayRuleName,omitempty"`
	// example:
	//
	// 100
	GrayRulePriority *string `json:"GrayRulePriority,omitempty" xml:"GrayRulePriority,omitempty"`
	// example:
	//
	// Tags
	GrayType *string `json:"GrayType,omitempty" xml:"GrayType,omitempty"`
	// example:
	//
	// dubbo
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// 1742277568000
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// example:
	//
	// 87d14faf58a103ac8840b9c5f1c2a0fe
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s QueryNacosGrayConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryNacosGrayConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNacosGrayConfigResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *QueryNacosGrayConfigResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *QueryNacosGrayConfigResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *QueryNacosGrayConfigResponseBodyData) GetGrayRule() *string {
	return s.GrayRule
}

func (s *QueryNacosGrayConfigResponseBodyData) GetGrayRuleName() *string {
	return s.GrayRuleName
}

func (s *QueryNacosGrayConfigResponseBodyData) GetGrayRulePriority() *string {
	return s.GrayRulePriority
}

func (s *QueryNacosGrayConfigResponseBodyData) GetGrayType() *string {
	return s.GrayType
}

func (s *QueryNacosGrayConfigResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *QueryNacosGrayConfigResponseBodyData) GetLastModified() *string {
	return s.LastModified
}

func (s *QueryNacosGrayConfigResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *QueryNacosGrayConfigResponseBodyData) SetAppName(v string) *QueryNacosGrayConfigResponseBodyData {
	s.AppName = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetContent(v string) *QueryNacosGrayConfigResponseBodyData {
	s.Content = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetDataId(v string) *QueryNacosGrayConfigResponseBodyData {
	s.DataId = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetGrayRule(v string) *QueryNacosGrayConfigResponseBodyData {
	s.GrayRule = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetGrayRuleName(v string) *QueryNacosGrayConfigResponseBodyData {
	s.GrayRuleName = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetGrayRulePriority(v string) *QueryNacosGrayConfigResponseBodyData {
	s.GrayRulePriority = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetGrayType(v string) *QueryNacosGrayConfigResponseBodyData {
	s.GrayType = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetGroup(v string) *QueryNacosGrayConfigResponseBodyData {
	s.Group = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetLastModified(v string) *QueryNacosGrayConfigResponseBodyData {
	s.LastModified = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) SetMd5(v string) *QueryNacosGrayConfigResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *QueryNacosGrayConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
