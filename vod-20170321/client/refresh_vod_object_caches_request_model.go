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
	// Specifies whether to purge resources in a directory if the resources requested are different from the resources on the origin server.
	//
	// 	- **true**: refreshes all resources in the directory. If you set this parameter to true, when the requested content matches the resource in the directory, the POP retrieves the resource from the origin server, returns the resource to the client, and caches the resource.
	//
	// 	- **false*	- (default): refreshes the changed resources in the directory. If you set this parameter to false, when the requested content matches the resource in the directory, the POP obtains the Last-Modified parameter of the resource from the origin server. If the value of the obtained Last-Modified parameter is the same as that of the cached resource, the cached resource is returned. Otherwise, the POP retrieves the resource from the origin server, returns the resource to the client, and caches the resource.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The URL of the file to be prefetched. Separate multiple URLs with line breaks (\\n or \\r\\n).
	//
	// This parameter is required.
	//
	// example:
	//
	// abc.com/image/1.png
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The type of the object that you want to refresh. Valid values:
	//
	// 	- **File*	- (default): refreshes one or more files.
	//
	// 	- **Directory**: refreshes the files in specified directories.
	//
	// 	- **Regex**: refreshes content based on regular expressions.
	//
	// 	- **IgnoreParams**: removes the question mark (?) and parameters after the question mark (?) in a request URL and refreshes content. After you call this operation with the request URL submitted, the system compares the submitted URL with the URL of the cached resource without specific parameters. If the URLs match, the POPs refresh the cached resource.
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
