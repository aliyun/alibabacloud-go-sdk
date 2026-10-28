// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionIds(v []*string) *DeleteAlertActionsRequest
	GetAlertActionIds() []*string
}

type DeleteAlertActionsRequest struct {
	// The unique identifiers of the alert action integrations to delete.
	//
	// This parameter is required.
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
}

func (s DeleteAlertActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertActionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertActionsRequest) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *DeleteAlertActionsRequest) SetAlertActionIds(v []*string) *DeleteAlertActionsRequest {
	s.AlertActionIds = v
	return s
}

func (s *DeleteAlertActionsRequest) Validate() error {
	return dara.Validate(s)
}
