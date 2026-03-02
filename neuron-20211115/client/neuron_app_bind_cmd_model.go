// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppBindCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *NeuronAppBindCmd
	GetAppId() *int64
}

type NeuronAppBindCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s NeuronAppBindCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppBindCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppBindCmd) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppBindCmd) SetAppId(v int64) *NeuronAppBindCmd {
	s.AppId = &v
	return s
}

func (s *NeuronAppBindCmd) Validate() error {
	return dara.Validate(s)
}
