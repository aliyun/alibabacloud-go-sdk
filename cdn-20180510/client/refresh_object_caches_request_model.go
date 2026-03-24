// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RefreshObjectCachesRequest
	GetForce() *bool
	SetObjectPath(v string) *RefreshObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshObjectCachesRequest
	GetObjectType() *string
	SetOwnerId(v int64) *RefreshObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *RefreshObjectCachesRequest
	GetSecurityToken() *string
}

type RefreshObjectCachesRequest struct {
	// When the comparison between the source content and the source site resources is consistent, should the resources within the corresponding range be forcibly refreshed. The default is false.
	//
	// 	- **true**: purges all resources in the range that corresponds to the type of the purge task. If you set this parameter to true, when the requested resource matches the resource in the range that corresponds to the type of the purge task, the POP retrieves the resource from the origin server, returns the resource to the client, and caches the resource.
	//
	// 	- **false**: purges the changed resources in the range that corresponds to the type of the purge task. If you set this parameter to false, when the requested resource matches the resource in the range that corresponds to the type of the purge task, the POP obtains the Last-Modified parameter of the resource from the origin server. If the obtained value of the Last-Modified parameter is the same as that of the cached resource, the cached resource is returned. Otherwise, the POP retrieves the resource from the origin server, returns the resource to the client, and caches the resource.
	//
	// >  This parameter takes effect only when the ObjectType parameter is not set to File.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// 	- If you submit multiple URLs or directories at a time, separate them with line breaks (\\n) or (\\r\\n).
	//
	// 	- The total number of domain names contained all URLs in a submitted task cannot exceed 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/image/1.png\\nhttp://aliyundoc.com/image/2.png
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the object that you want to refresh. Valid values:
	//
	// 	- **File*	- (default): refreshes one or more files.
	//
	// 	- **Directory**: refreshes the files in one or more directories.
	//
	// 	- **Regex**: refreshes content based on regular expressions.
	//
	// 	- **ExQuery**: omits parameters after the question mark in the URL and refreshes content.
	//
	// If you set the ObjectType parameter to File or Directory, you can view [Refresh and prefetch resources](https://help.aliyun.com/document_detail/27140.html) to obtain more information. If you set the ObjectType parameter to Regex, you can view [Configure URL refresh rules that contain regular expressions](https://help.aliyun.com/document_detail/146195.html) to obtain more information.
	//
	// If you set the ObjectType parameter to Directory, the resources in the directory that you want to refresh are marked as expired. You cannot delete the directory. If clients request resources on POPs that are marked as expired, Alibaba Cloud CDN checks whether the resources on your origin server are updated. If resources are updated, Alibaba Cloud CDN retrieves the latest version of the resources and returns the resources to the clients. Otherwise, the origin server returns the 304 status code.
	//
	// example:
	//
	// File
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RefreshObjectCachesRequest) SetForce(v bool) *RefreshObjectCachesRequest {
	s.Force = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectPath(v string) *RefreshObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectType(v string) *RefreshObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetOwnerId(v int64) *RefreshObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetSecurityToken(v string) *RefreshObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RefreshObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
