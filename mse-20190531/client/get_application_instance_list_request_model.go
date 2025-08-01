// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetApplicationInstanceListRequest
	GetAcceptLanguage() *string
	SetAppId(v string) *GetApplicationInstanceListRequest
	GetAppId() *string
	SetAppName(v string) *GetApplicationInstanceListRequest
	GetAppName() *string
	SetNamespace(v string) *GetApplicationInstanceListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *GetApplicationInstanceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetApplicationInstanceListRequest
	GetPageSize() *int32
	SetRegion(v string) *GetApplicationInstanceListRequest
	GetRegion() *string
	SetTag(v string) *GetApplicationInstanceListRequest
	GetTag() *string
}

type GetApplicationInstanceListRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The application ID.
	//
	// example:
	//
	// abcde@12345
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The Microservices Engine (MSE) namespace to which the application belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The tags that you want to add to nodes.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetApplicationInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationInstanceListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetApplicationInstanceListRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationInstanceListRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationInstanceListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetApplicationInstanceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetApplicationInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetApplicationInstanceListRequest) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationInstanceListRequest) GetTag() *string {
	return s.Tag
}

func (s *GetApplicationInstanceListRequest) SetAcceptLanguage(v string) *GetApplicationInstanceListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetAppId(v string) *GetApplicationInstanceListRequest {
	s.AppId = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetAppName(v string) *GetApplicationInstanceListRequest {
	s.AppName = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetNamespace(v string) *GetApplicationInstanceListRequest {
	s.Namespace = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetPageNumber(v int32) *GetApplicationInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetPageSize(v int32) *GetApplicationInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetRegion(v string) *GetApplicationInstanceListRequest {
	s.Region = &v
	return s
}

func (s *GetApplicationInstanceListRequest) SetTag(v string) *GetApplicationInstanceListRequest {
	s.Tag = &v
	return s
}

func (s *GetApplicationInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
