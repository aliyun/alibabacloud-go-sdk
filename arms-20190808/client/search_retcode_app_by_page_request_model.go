// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRetcodeAppByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *SearchRetcodeAppByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRetcodeAppByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchRetcodeAppByPageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SearchRetcodeAppByPageRequest
	GetResourceGroupId() *string
	SetRetcodeAppId(v string) *SearchRetcodeAppByPageRequest
	GetRetcodeAppId() *string
	SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest
	GetRetcodeAppName() *string
	SetTags(v []*SearchRetcodeAppByPageRequestTags) *SearchRetcodeAppByPageRequest
	GetTags() []*SearchRetcodeAppByPageRequestTags
}

type SearchRetcodeAppByPageRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain the resource group ID in the **Resource Management*	- console.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The application ID.
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of an application. The URL in the browser address bar contains the pid of this application in the format of `pid=xxx`. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is `xxx%4074xxx`, you must replace **%40*	- with the at sign (@). The actual PID is `xxx@74xxx`.
	//
	// example:
	//
	// eb4zdose6v@9781be0f44d****
	RetcodeAppId *string `json:"RetcodeAppId,omitempty" xml:"RetcodeAppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// App1
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	// The tag.
	Tags []*SearchRetcodeAppByPageRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s SearchRetcodeAppByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRetcodeAppByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRetcodeAppByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchRetcodeAppByPageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchRetcodeAppByPageRequest) GetRetcodeAppId() *string {
	return s.RetcodeAppId
}

func (s *SearchRetcodeAppByPageRequest) GetRetcodeAppName() *string {
	return s.RetcodeAppName
}

func (s *SearchRetcodeAppByPageRequest) GetTags() []*SearchRetcodeAppByPageRequestTags {
	return s.Tags
}

func (s *SearchRetcodeAppByPageRequest) SetPageNumber(v int32) *SearchRetcodeAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetPageSize(v int32) *SearchRetcodeAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRegionId(v string) *SearchRetcodeAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetResourceGroupId(v string) *SearchRetcodeAppByPageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppId(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetTags(v []*SearchRetcodeAppByPageRequestTags) *SearchRetcodeAppByPageRequest {
	s.Tags = v
	return s
}

func (s *SearchRetcodeAppByPageRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchRetcodeAppByPageRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchRetcodeAppByPageRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageRequestTags) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequestTags) GetKey() *string {
	return s.Key
}

func (s *SearchRetcodeAppByPageRequestTags) GetValue() *string {
	return s.Value
}

func (s *SearchRetcodeAppByPageRequestTags) SetKey(v string) *SearchRetcodeAppByPageRequestTags {
	s.Key = &v
	return s
}

func (s *SearchRetcodeAppByPageRequestTags) SetValue(v string) *SearchRetcodeAppByPageRequestTags {
	s.Value = &v
	return s
}

func (s *SearchRetcodeAppByPageRequestTags) Validate() error {
	return dara.Validate(s)
}
