// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshVodObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RefreshVodObjectCachesRequest
	GetForce() *bool
	SetObjectPath(v string) *RefreshVodObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshVodObjectCachesRequest
	GetObjectType() *string
	SetOwnerId(v int64) *RefreshVodObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *RefreshVodObjectCachesRequest
	GetSecurityToken() *string
}

type RefreshVodObjectCachesRequest struct {
	// Specifies whether to purge all resources in the corresponding directory when the back-to-origin content is inconsistent with the origin server resources. Default value: false.
	//
	// - **true**: purges all resources in the corresponding directory. When "Purge All Resources" is selected, if the requested content matches a resource in the directory, the CDN node fetches the new resource from the origin server, returns it to the user, and re-caches the resource.
	//
	// - **false**: purges only changed resources in the corresponding directory. When "Purge Changed Resources" is selected, if the requested content matches a resource in the directory, the CDN node retrieves the Last-Modified information of the resource from the origin server. If it matches the currently cached resource, the cached resource is returned directly. If it does not match, the CDN node fetches the new resource from the origin server, returns it to the user, and re-caches the resource.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The URL of the file to prefetch. Separate multiple URLs with line breaks (
	//
	//  or
	//
	// ).
	//
	// This parameter is required.
	//
	// example:
	//
	// abc.com/image/1.png
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of purge. Valid values:
	//
	// - **File*	- (default): file purge.
	//
	// - **Directory**: directory purge.
	//
	// - **Regex**: regular expression-based purge.
	//
	// - **IgnoreParams**: parameter-stripped purge. Parameter stripping refers to removing the question mark (?) and all characters after it from the request URL. Parameter-stripped purge means that you submit a parameter-stripped URL through the API, and the submitted URL is matched against cached resource URLs after parameter stripping. If a cached resource URL matches the submitted URL after parameter stripping, the CDN node purges the cached resource.
	//
	// example:
	//
	// File
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshVodObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshVodObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshVodObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshVodObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshVodObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshVodObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RefreshVodObjectCachesRequest) SetForce(v bool) *RefreshVodObjectCachesRequest {
	s.Force = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetObjectPath(v string) *RefreshVodObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetObjectType(v string) *RefreshVodObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetOwnerId(v int64) *RefreshVodObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetSecurityToken(v string) *RefreshVodObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
