// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertNotifyTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateIds(v []*string) *DeleteAlertNotifyTemplatesRequest
	GetAlertNotifyTemplateIds() []*string
}

type DeleteAlertNotifyTemplatesRequest struct {
	// The IDs of the notification templates.
	//
	// This parameter is required.
	AlertNotifyTemplateIds []*string `json:"alertNotifyTemplateIds,omitempty" xml:"alertNotifyTemplateIds,omitempty" type:"Repeated"`
}

func (s DeleteAlertNotifyTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertNotifyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertNotifyTemplatesRequest) GetAlertNotifyTemplateIds() []*string {
	return s.AlertNotifyTemplateIds
}

func (s *DeleteAlertNotifyTemplatesRequest) SetAlertNotifyTemplateIds(v []*string) *DeleteAlertNotifyTemplatesRequest {
	s.AlertNotifyTemplateIds = v
	return s
}

func (s *DeleteAlertNotifyTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
