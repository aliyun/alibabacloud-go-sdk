// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationEcuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListApplicationEcuRequest
	GetAppId() *string
	SetLogicalRegionId(v string) *ListApplicationEcuRequest
	GetLogicalRegionId() *string
}

type ListApplicationEcuRequest struct {
	// The ID of the application whose ECUs you want to query. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// e809****-43d7-4c6b-8e01-b0d9d1db****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the microservices namespace.
	//
	// example:
	//
	// cn-hangzhou:***wei
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s ListApplicationEcuRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationEcuRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationEcuRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListApplicationEcuRequest) SetAppId(v string) *ListApplicationEcuRequest {
	s.AppId = &v
	return s
}

func (s *ListApplicationEcuRequest) SetLogicalRegionId(v string) *ListApplicationEcuRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListApplicationEcuRequest) Validate() error {
	return dara.Validate(s)
}
