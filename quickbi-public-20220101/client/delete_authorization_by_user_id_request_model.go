// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQbiUserId(v string) *DeleteAuthorizationByUserIdRequest
	GetQbiUserId() *string
	SetResourceId(v string) *DeleteAuthorizationByUserIdRequest
	GetResourceId() *string
	SetResourceType(v string) *DeleteAuthorizationByUserIdRequest
	GetResourceType() *string
}

type DeleteAuthorizationByUserIdRequest struct {
	// The Quick BI user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc4b***94fa4
	QbiUserId *string `json:"QbiUserId,omitempty" xml:"QbiUserId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - dashboard: dashboard
	//
	// - report: workbook
	//
	// - dashboardOfflineQuery: self-service data retrieval
	//
	// - cube: dataset
	//
	// - datasource: data source
	//
	// - screen: data dashboard
	//
	// - ANALYSIS: ad hoc analysis
	//
	// - dataForm: data entry form
	//
	// This parameter is required.
	//
	// example:
	//
	// report
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteAuthorizationByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationByUserIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationByUserIdRequest) GetQbiUserId() *string {
	return s.QbiUserId
}

func (s *DeleteAuthorizationByUserIdRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DeleteAuthorizationByUserIdRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteAuthorizationByUserIdRequest) SetQbiUserId(v string) *DeleteAuthorizationByUserIdRequest {
	s.QbiUserId = &v
	return s
}

func (s *DeleteAuthorizationByUserIdRequest) SetResourceId(v string) *DeleteAuthorizationByUserIdRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteAuthorizationByUserIdRequest) SetResourceType(v string) *DeleteAuthorizationByUserIdRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteAuthorizationByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
