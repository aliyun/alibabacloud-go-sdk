// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIImageInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIImageInfoIds(v string) *DeleteAIImageInfosRequest
	GetAIImageInfoIds() *string
}

type DeleteAIImageInfosRequest struct {
	// The IDs of AI image information. This parameter consists of one or more AIImageInfoId values. The AIImageInfoId is the value of the AIImageInfoId response parameter returned by the [ListAIImageInfo](https://help.aliyun.com/document_detail/186924.html) operation.
	//
	// - A maximum of 10 IDs are supported.
	//
	// - Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// b89a6aabf144*****6197ebd6fe6cf29
	AIImageInfoIds *string `json:"AIImageInfoIds,omitempty" xml:"AIImageInfoIds,omitempty"`
}

func (s DeleteAIImageInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIImageInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosRequest) GetAIImageInfoIds() *string {
	return s.AIImageInfoIds
}

func (s *DeleteAIImageInfosRequest) SetAIImageInfoIds(v string) *DeleteAIImageInfosRequest {
	s.AIImageInfoIds = &v
	return s
}

func (s *DeleteAIImageInfosRequest) Validate() error {
	return dara.Validate(s)
}
