// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChainId(v string) *DeleteChainRequest
	GetChainId() *string
	SetInstanceId(v string) *DeleteChainRequest
	GetInstanceId() *string
}

type DeleteChainRequest struct {
	// The ID of the delivery pipeline.
	//
	// This parameter is required.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteChainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChainRequest) GoString() string {
	return s.String()
}

func (s *DeleteChainRequest) GetChainId() *string {
	return s.ChainId
}

func (s *DeleteChainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteChainRequest) SetChainId(v string) *DeleteChainRequest {
	s.ChainId = &v
	return s
}

func (s *DeleteChainRequest) SetInstanceId(v string) *DeleteChainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChainRequest) Validate() error {
	return dara.Validate(s)
}
