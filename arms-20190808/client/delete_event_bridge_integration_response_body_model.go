// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBridgeIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteEventBridgeIntegrationResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteEventBridgeIntegrationResponseBody
	GetRequestId() *string
}

type DeleteEventBridgeIntegrationResponseBody struct {
	// Indicates whether the EventBridge integration is deleted.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2B289756-E791-5842-BCBD-AD414C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventBridgeIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBridgeIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteEventBridgeIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventBridgeIntegrationResponseBody) SetIsSuccess(v bool) *DeleteEventBridgeIntegrationResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteEventBridgeIntegrationResponseBody) SetRequestId(v string) *DeleteEventBridgeIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventBridgeIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
