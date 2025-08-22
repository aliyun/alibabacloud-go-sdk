// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RefreshErObjectCachesRequest
	GetForce() *bool
	SetMergeDomainName(v string) *RefreshErObjectCachesRequest
	GetMergeDomainName() *string
	SetObjectPath(v string) *RefreshErObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshErObjectCachesRequest
	GetObjectType() *string
	SetRoutineId(v string) *RefreshErObjectCachesRequest
	GetRoutineId() *string
}

type RefreshErObjectCachesRequest struct {
	// Specifies whether to refresh resources in a directory if the resources requested are different from the resources on the origin server. Default value: false.
	//
	// 	- **true**: refreshes all resources in the directory.
	//
	// 	- **false**: refreshes the changed resources in the directory.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The domain names that are merged for refreshing. POPs that provide services for the domain names are refreshed.
	//
	// >  Separate multiple domain names with commas (,).
	//
	// example:
	//
	// a.test.com,b.test.com
	MergeDomainName *string `json:"MergeDomainName,omitempty" xml:"MergeDomainName,omitempty"`
	// The URL that you want to refresh.
	//
	// > 	- Separate URLs with line feeds (\\n or \\r\\n). Each object path can be up to 1,024 characters in length.
	//
	// >	- The URLs in a request must belong to the same domain name.
	//
	// >	- You can refresh up to 1,000 URLs in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/examplefile.txt
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	// The refresh type. Valid values:
	//
	// 	- **File*	- (default): refreshes content based on URLs.
	//
	// 	- **Directory**: refreshes content based on directories.
	//
	// 	- **Regex**: refreshes content based on regular expressions.
	//
	// 	- **IgnoreParams**: removes the question mark (`?`) and parameters after the question mark (`?`) in a request URL and refreshes content. After you call this operation with the request URL submitted, the system compares the submitted URL with the URL of the cached resource without specific parameters. If the URLs match, the points of presence (POPs) refresh the cached resource.
	//
	// >  If you refresh the files in one or more directories, the resources in the directory that you want to refresh are marked as expired. You cannot delete the directory. If clients request resources on POPs that are marked as expired, Dynamic Content Delivery Network (DCDN) checks whether the resources on your origin server are updated. If resources are updated, DCDN retrieves the latest version of the resources and returns the resources to the clients. Otherwise, the origin server returns the 304 status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// File
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the routine, which is in the format of "Name.Subdomain" and is the unique identifier of a custom routine.
	//
	// example:
	//
	// test.mysubdomain
	RoutineId *string `json:"RoutineId,omitempty" xml:"RoutineId,omitempty"`
}

func (s RefreshErObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshErObjectCachesRequest) GetMergeDomainName() *string {
	return s.MergeDomainName
}

func (s *RefreshErObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshErObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshErObjectCachesRequest) GetRoutineId() *string {
	return s.RoutineId
}

func (s *RefreshErObjectCachesRequest) SetForce(v bool) *RefreshErObjectCachesRequest {
	s.Force = &v
	return s
}

func (s *RefreshErObjectCachesRequest) SetMergeDomainName(v string) *RefreshErObjectCachesRequest {
	s.MergeDomainName = &v
	return s
}

func (s *RefreshErObjectCachesRequest) SetObjectPath(v string) *RefreshErObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshErObjectCachesRequest) SetObjectType(v string) *RefreshErObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshErObjectCachesRequest) SetRoutineId(v string) *RefreshErObjectCachesRequest {
	s.RoutineId = &v
	return s
}

func (s *RefreshErObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
