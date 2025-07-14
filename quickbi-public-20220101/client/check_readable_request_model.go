// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReadableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *CheckReadableRequest
	GetUserId() *string
	SetWorksId(v string) *CheckReadableRequest
	GetWorksId() *string
}

type CheckReadableRequest struct {
	// The user ID of the Quick BI to be checked.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the work. Resources here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CheckReadableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckReadableRequest) GoString() string {
	return s.String()
}

func (s *CheckReadableRequest) GetUserId() *string {
	return s.UserId
}

func (s *CheckReadableRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *CheckReadableRequest) SetUserId(v string) *CheckReadableRequest {
	s.UserId = &v
	return s
}

func (s *CheckReadableRequest) SetWorksId(v string) *CheckReadableRequest {
	s.WorksId = &v
	return s
}

func (s *CheckReadableRequest) Validate() error {
	return dara.Validate(s)
}
