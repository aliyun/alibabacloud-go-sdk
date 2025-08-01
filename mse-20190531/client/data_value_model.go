// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *DataValue
	GetUserId() *string
	SetRegionId(v string) *DataValue
	GetRegionId() *string
	SetNamespace(v string) *DataValue
	GetNamespace() *string
	SetAppName(v string) *DataValue
	GetAppName() *string
	SetAppId(v string) *DataValue
	GetAppId() *string
}

type DataValue struct {
	// The ID of the user to which the application belongs.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The region where the application resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The microservice namespace where the application resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The application name.
	//
	// example:
	//
	// example-app-name
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 123456abcde@12345abcde
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetUserId() *string {
	return s.UserId
}

func (s *DataValue) GetRegionId() *string {
	return s.RegionId
}

func (s *DataValue) GetNamespace() *string {
	return s.Namespace
}

func (s *DataValue) GetAppName() *string {
	return s.AppName
}

func (s *DataValue) GetAppId() *string {
	return s.AppId
}

func (s *DataValue) SetUserId(v string) *DataValue {
	s.UserId = &v
	return s
}

func (s *DataValue) SetRegionId(v string) *DataValue {
	s.RegionId = &v
	return s
}

func (s *DataValue) SetNamespace(v string) *DataValue {
	s.Namespace = &v
	return s
}

func (s *DataValue) SetAppName(v string) *DataValue {
	s.AppName = &v
	return s
}

func (s *DataValue) SetAppId(v string) *DataValue {
	s.AppId = &v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
