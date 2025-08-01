// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v *GetNacosConfigResponseBodyConfiguration) *GetNacosConfigResponseBody
	GetConfiguration() *GetNacosConfigResponseBodyConfiguration
	SetErrorCode(v string) *GetNacosConfigResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetNacosConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNacosConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNacosConfigResponseBody
	GetSuccess() *bool
}

type GetNacosConfigResponseBody struct {
	// Configuration information.
	Configuration *GetNacosConfigResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4EAB48C-BB4B-5B8D-B33B-35D69606C5AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request, with values as follows:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNacosConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponseBody) GetConfiguration() *GetNacosConfigResponseBodyConfiguration {
	return s.Configuration
}

func (s *GetNacosConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNacosConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNacosConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNacosConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNacosConfigResponseBody) SetConfiguration(v *GetNacosConfigResponseBodyConfiguration) *GetNacosConfigResponseBody {
	s.Configuration = v
	return s
}

func (s *GetNacosConfigResponseBody) SetErrorCode(v string) *GetNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetMessage(v string) *GetNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetRequestId(v string) *GetNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetSuccess(v bool) *GetNacosConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetNacosConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNacosConfigResponseBodyConfiguration struct {
	// Application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// List of IPs for Beta release.
	//
	// example:
	//
	// 1.1.XX.XX，2.2.XX.XX
	BetaIps *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	// Configuration content.
	//
	// example:
	//
	// log.level=error
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// log.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Configuration description.
	//
	// example:
	//
	// For testing
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Encrypted key.
	//
	// example:
	//
	// key
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	// Current gray version information
	GrayVersions []*GetNacosConfigResponseBodyConfigurationGrayVersions `json:"GrayVersions,omitempty" xml:"GrayVersions,omitempty" type:"Repeated"`
	// Configuration group name.
	//
	// example:
	//
	// test
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// Message digest of the configuration.
	//
	// example:
	//
	// 123rfsdf3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// Tags of the configuration.
	//
	// example:
	//
	// context
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Format of the configuration content.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNacosConfigResponseBodyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetNacosConfigResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponseBodyConfiguration) GetAppName() *string {
	return s.AppName
}

func (s *GetNacosConfigResponseBodyConfiguration) GetBetaIps() *string {
	return s.BetaIps
}

func (s *GetNacosConfigResponseBodyConfiguration) GetContent() *string {
	return s.Content
}

func (s *GetNacosConfigResponseBodyConfiguration) GetDataId() *string {
	return s.DataId
}

func (s *GetNacosConfigResponseBodyConfiguration) GetDesc() *string {
	return s.Desc
}

func (s *GetNacosConfigResponseBodyConfiguration) GetEncryptedDataKey() *string {
	return s.EncryptedDataKey
}

func (s *GetNacosConfigResponseBodyConfiguration) GetGrayVersions() []*GetNacosConfigResponseBodyConfigurationGrayVersions {
	return s.GrayVersions
}

func (s *GetNacosConfigResponseBodyConfiguration) GetGroup() *string {
	return s.Group
}

func (s *GetNacosConfigResponseBodyConfiguration) GetMd5() *string {
	return s.Md5
}

func (s *GetNacosConfigResponseBodyConfiguration) GetTags() *string {
	return s.Tags
}

func (s *GetNacosConfigResponseBodyConfiguration) GetType() *string {
	return s.Type
}

func (s *GetNacosConfigResponseBodyConfiguration) SetAppName(v string) *GetNacosConfigResponseBodyConfiguration {
	s.AppName = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetBetaIps(v string) *GetNacosConfigResponseBodyConfiguration {
	s.BetaIps = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetContent(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Content = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetDataId(v string) *GetNacosConfigResponseBodyConfiguration {
	s.DataId = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetDesc(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Desc = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetEncryptedDataKey(v string) *GetNacosConfigResponseBodyConfiguration {
	s.EncryptedDataKey = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetGrayVersions(v []*GetNacosConfigResponseBodyConfigurationGrayVersions) *GetNacosConfigResponseBodyConfiguration {
	s.GrayVersions = v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetGroup(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Group = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetMd5(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Md5 = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetTags(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Tags = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetType(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Type = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetNacosConfigResponseBodyConfigurationGrayVersions struct {
	// Gray version name
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the current gray rule.
	//
	// example:
	//
	// 20
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Rules of the current gray version
	//
	// example:
	//
	// a=b
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Gray type
	//
	// example:
	//
	// Beta
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNacosConfigResponseBodyConfigurationGrayVersions) String() string {
	return dara.Prettify(s)
}

func (s GetNacosConfigResponseBodyConfigurationGrayVersions) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) GetName() *string {
	return s.Name
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) GetPriority() *int32 {
	return s.Priority
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) GetRule() *string {
	return s.Rule
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) GetType() *string {
	return s.Type
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) SetName(v string) *GetNacosConfigResponseBodyConfigurationGrayVersions {
	s.Name = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) SetPriority(v int32) *GetNacosConfigResponseBodyConfigurationGrayVersions {
	s.Priority = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) SetRule(v string) *GetNacosConfigResponseBodyConfigurationGrayVersions {
	s.Rule = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) SetType(v string) *GetNacosConfigResponseBodyConfigurationGrayVersions {
	s.Type = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfigurationGrayVersions) Validate() error {
	return dara.Validate(s)
}
