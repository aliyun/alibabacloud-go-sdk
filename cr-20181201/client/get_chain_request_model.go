// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChainId(v string) *GetChainRequest
	GetChainId() *string
	SetInstanceId(v string) *GetChainRequest
	GetInstanceId() *string
}

type GetChainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetChainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChainRequest) GoString() string {
	return s.String()
}

func (s *GetChainRequest) GetChainId() *string {
	return s.ChainId
}

func (s *GetChainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChainRequest) SetChainId(v string) *GetChainRequest {
	s.ChainId = &v
	return s
}

func (s *GetChainRequest) SetInstanceId(v string) *GetChainRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChainRequest) Validate() error {
	return dara.Validate(s)
}
