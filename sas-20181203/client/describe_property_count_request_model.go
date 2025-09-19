// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribePropertyCountRequest
	GetType() *string
	SetUuidList(v string) *DescribePropertyCountRequest
	GetUuidList() *string
}

type DescribePropertyCountRequest struct {
	// The type of the asset fingerprints. Separate multiple types with commas (,). Valid values:
	//
	// 	- **port**: port
	//
	// 	- **process**: process
	//
	// 	- **software**: software
	//
	// 	- **user**: account
	//
	// 	- **cron**: scheduled task
	//
	// 	- **sca**: middleware
	//
	// 	- **web**: website
	//
	// 	- **database**: database
	//
	// 	- **lkm**: kernel module
	//
	// 	- **autorun**: startup item
	//
	// 	- **web_server**: web service
	//
	// example:
	//
	// port,process
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the assets. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 9658314a-7609-4426-afc4-2c924072****,
	//
	// 9658314a-7609-4426-afc4-2c924072****
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s DescribePropertyCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountRequest) GetType() *string {
	return s.Type
}

func (s *DescribePropertyCountRequest) GetUuidList() *string {
	return s.UuidList
}

func (s *DescribePropertyCountRequest) SetType(v string) *DescribePropertyCountRequest {
	s.Type = &v
	return s
}

func (s *DescribePropertyCountRequest) SetUuidList(v string) *DescribePropertyCountRequest {
	s.UuidList = &v
	return s
}

func (s *DescribePropertyCountRequest) Validate() error {
	return dara.Validate(s)
}
