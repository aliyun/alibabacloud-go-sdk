// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyQuotaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateApiKeyQuotaShrinkRequest
	GetInstanceId() *string
	SetKeysShrink(v string) *UpdateApiKeyQuotaShrinkRequest
	GetKeysShrink() *string
}

type UpdateApiKeyQuotaShrinkRequest struct {
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s UpdateApiKeyQuotaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyQuotaShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyQuotaShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApiKeyQuotaShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *UpdateApiKeyQuotaShrinkRequest) SetInstanceId(v string) *UpdateApiKeyQuotaShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApiKeyQuotaShrinkRequest) SetKeysShrink(v string) *UpdateApiKeyQuotaShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *UpdateApiKeyQuotaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
