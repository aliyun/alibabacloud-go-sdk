// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RefreshDcdnObjectCachesRequest
	GetForce() *bool
	SetObjectPath(v string) *RefreshDcdnObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshDcdnObjectCachesRequest
	GetObjectType() *string
	SetOwnerId(v int64) *RefreshDcdnObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *RefreshDcdnObjectCachesRequest
	GetSecurityToken() *string
}

type RefreshDcdnObjectCachesRequest struct {
	// Specifies whether to refresh resources in a directory if the resources are different from the resources in the same directory in the origin server. Default value: false.
	//
	// 	- **true**: refresh all resources in the directory.
	//
	// 	- **false**: refresh the changed resources in the directory.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The path of the objects that you want to refresh. Separate multiple URLs with line feed characters (\\n) or a pair of carriage return and line feed characters (\\r\\n).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com/example.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The refresh type. Valid values:
	//
	// 	- **File*	- (default): refreshes resources based on URLs.
	//
	// 	- **Directory**: refreshes resources based on directories.
	//
	// 	- **Regex**: refreshes content based on regular expressions.
	//
	// 	- **IgnoreParams**: removes the question mark (`?`) and parameters after `?` in a request URL and refreshes content. After you call this operation with the request URL submitted, the system compares the submitted URL with the URL of the cached resource without specific parameters. If the URLs match, the DCDN POPs refresh the cached resource.
	//
	// >	- For more information about features of URL refresh and directory refresh, see [Refresh and prefetch resources](https://help.aliyun.com/document_detail/64936.html).
	//
	// >	- If you set ObjectType to Directory, the resources in the directory that you want to refresh are marked as expired. You cannot delete the directory. If clients request resources after the resources on POPs are marked as expired, DCDN checks whether the resources on your origin server are updated with a later version. If a later version exists, DCDN retrieves the resources of the later version and returns the resources to the clients. Otherwise, DCDN retrieves the 304 status code from the origin server.
	//
	// example:
	//
	// File
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshDcdnObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshDcdnObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshDcdnObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshDcdnObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshDcdnObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RefreshDcdnObjectCachesRequest) SetForce(v bool) *RefreshDcdnObjectCachesRequest {
	s.Force = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetObjectPath(v string) *RefreshDcdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetObjectType(v string) *RefreshDcdnObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetOwnerId(v int64) *RefreshDcdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) SetSecurityToken(v string) *RefreshDcdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RefreshDcdnObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
