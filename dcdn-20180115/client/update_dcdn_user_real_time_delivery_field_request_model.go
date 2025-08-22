// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnUserRealTimeDeliveryFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest
	GetBusinessType() *string
	SetFields(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest
	GetFields() *string
}

type UpdateDcdnUserRealTimeDeliveryFieldRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of L1 Dynamic Route for CDN (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The list of fields. Separate multiple fields with commas (,). For more information, see [Fields in a real-time log](https://help.aliyun.com/document_detail/324199.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// body_bytes_sent,client_ip,content_type
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) GetFields() *string {
	return s.Fields
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) SetBusinessType(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) SetFields(v string) *UpdateDcdnUserRealTimeDeliveryFieldRequest {
	s.Fields = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldRequest) Validate() error {
	return dara.Validate(s)
}
