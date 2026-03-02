// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppVersionOpCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersionId(v int64) *NeuronAppVersionOpCmd
	GetAppVersionId() *int64
}

type NeuronAppVersionOpCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppVersionId *int64 `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
}

func (s NeuronAppVersionOpCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppVersionOpCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppVersionOpCmd) GetAppVersionId() *int64 {
	return s.AppVersionId
}

func (s *NeuronAppVersionOpCmd) SetAppVersionId(v int64) *NeuronAppVersionOpCmd {
	s.AppVersionId = &v
	return s
}

func (s *NeuronAppVersionOpCmd) Validate() error {
	return dara.Validate(s)
}
