// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronMobiTokenCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *NeuronMobiTokenCreateCmd
	GetAppId() *int64
}

type NeuronMobiTokenCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 233131
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s NeuronMobiTokenCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronMobiTokenCreateCmd) GoString() string {
	return s.String()
}

func (s *NeuronMobiTokenCreateCmd) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronMobiTokenCreateCmd) SetAppId(v int64) *NeuronMobiTokenCreateCmd {
	s.AppId = &v
	return s
}

func (s *NeuronMobiTokenCreateCmd) Validate() error {
	return dara.Validate(s)
}
