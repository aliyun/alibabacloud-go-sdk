// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppUnBindCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *NeuronAppUnBindCmd
	GetAppId() *int64
	SetReason(v string) *NeuronAppUnBindCmd
	GetReason() *string
}

type NeuronAppUnBindCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 价格
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s NeuronAppUnBindCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppUnBindCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppUnBindCmd) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppUnBindCmd) GetReason() *string {
	return s.Reason
}

func (s *NeuronAppUnBindCmd) SetAppId(v int64) *NeuronAppUnBindCmd {
	s.AppId = &v
	return s
}

func (s *NeuronAppUnBindCmd) SetReason(v string) *NeuronAppUnBindCmd {
	s.Reason = &v
	return s
}

func (s *NeuronAppUnBindCmd) Validate() error {
	return dara.Validate(s)
}
