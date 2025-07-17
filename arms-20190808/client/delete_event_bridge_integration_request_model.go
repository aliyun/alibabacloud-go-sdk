// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBridgeIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteEventBridgeIntegrationRequest
	GetId() *int64
}

type DeleteEventBridgeIntegrationRequest struct {
	// Required. The ID of the EventBridge notification integration. You can call the **ListEventBridgeIntegrations*	- operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteEventBridgeIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBridgeIntegrationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteEventBridgeIntegrationRequest) SetId(v int64) *DeleteEventBridgeIntegrationRequest {
	s.Id = &v
	return s
}

func (s *DeleteEventBridgeIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
