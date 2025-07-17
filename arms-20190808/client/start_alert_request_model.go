// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v string) *StartAlertRequest
	GetAlertId() *string
	SetRegionId(v string) *StartAlertRequest
	GetRegionId() *string
}

type StartAlertRequest struct {
	// The ID of the alert rule. You can call the SearchAlertRules operation and view the `Id` parameter in the response. For more information, see [SearchAlertRules](https://help.aliyun.com/document_detail/175825.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1610***
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The ID of the region. Set the value to `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) GetAlertId() *string {
	return s.AlertId
}

func (s *StartAlertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartAlertRequest) SetAlertId(v string) *StartAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StartAlertRequest) SetRegionId(v string) *StartAlertRequest {
	s.RegionId = &v
	return s
}

func (s *StartAlertRequest) Validate() error {
	return dara.Validate(s)
}
