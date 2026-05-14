// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiInstanceId(v string) *DeleteModelResourceRequest
	GetAiInstanceId() *string
}

type DeleteModelResourceRequest struct {
	// example:
	//
	// hologram_aicombo_public_cn-yi34cxxxx
	AiInstanceId *string `json:"aiInstanceId,omitempty" xml:"aiInstanceId,omitempty"`
}

func (s DeleteModelResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelResourceRequest) GetAiInstanceId() *string {
	return s.AiInstanceId
}

func (s *DeleteModelResourceRequest) SetAiInstanceId(v string) *DeleteModelResourceRequest {
	s.AiInstanceId = &v
	return s
}

func (s *DeleteModelResourceRequest) Validate() error {
	return dara.Validate(s)
}
