// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetAlertEventRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetAlertEventRequest
	GetOpTenantId() *int64
}

type GetAlertEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetAlertEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventRequest) GoString() string {
	return s.String()
}

func (s *GetAlertEventRequest) GetId() *int64 {
	return s.Id
}

func (s *GetAlertEventRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAlertEventRequest) SetId(v int64) *GetAlertEventRequest {
	s.Id = &v
	return s
}

func (s *GetAlertEventRequest) SetOpTenantId(v int64) *GetAlertEventRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAlertEventRequest) Validate() error {
	return dara.Validate(s)
}
